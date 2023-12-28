// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mmf

import (
	"fmt"
	"log"
	"time"

	"open-match.dev/open-match/pkg/matchfunction"
	"open-match.dev/open-match/pkg/pb"
)

// This match function fetches all the Tickets for all the pools specified in
// the profile. It uses a configured number of tickets from each pool to generate
// a Match Proposal. It continues to generate proposals till one of the pools
// runs out of Tickets.
const (
	matchName              = "basic-matchfunction"
	ticketsPerPoolPerMatch = 4
)

// Run is this match function's implementation of the gRPC call defined in api/matchfunction.proto.
func (s *MatchFunctionService) Run(req *pb.RunRequest, stream pb.MatchFunction_RunServer) error {
	// Fetch tickets for the pools specified in the Match Profile.
	log.Printf("Generating proposals for function %v", req.GetProfile().GetName())

	poolTickets, err := matchfunction.QueryPools(stream.Context(), s.queryServiceClient, req.GetProfile().GetPools())
	if err != nil {
		log.Printf("Failed to query tickets for the given pools, got %s", err.Error())
		return err
	}

	// Generate proposals.
	proposals, err := makeMatches(req.GetProfile(), poolTickets)
	if err != nil {
		log.Printf("Failed to generate matches, got %s", err.Error())
		return err
	}

	log.Printf("Streaming %v proposals to Open Match", len(proposals))
	// Stream the generated proposals back to Open Match.
	for _, proposal := range proposals {
		if err := stream.Send(&pb.RunResponse{Proposal: proposal}); err != nil {
			log.Printf("Failed to stream proposals to Open Match, got %s", err.Error())
			return err
		}
	}

	return nil
}

func makeMatches(p *pb.MatchProfile, poolTickets map[string][]*pb.Ticket) ([]*pb.Match, error) {
	var matches []*pb.Match

	var user2Tickets = map[uint64]*pb.Ticket{}
	for _, tickets := range poolTickets {
		for _, v := range tickets {
			userId := uint64(v.SearchFields.DoubleArgs["userId"])
			user2Tickets[userId] = v
		}
	}

	var invalidLoopCount int32
	var count uint64
	for {
		invalidLoopCount++
		if invalidLoopCount > 10 {
			break
		}

		count++
		matchTickets := []*pb.Ticket{}

		for userId, ticket := range user2Tickets {
			if userId%2 == count%2 {
				matchTickets = append(matchTickets, ticket)
			}
		}

		if len(matchTickets) < ticketsPerPoolPerMatch {
			continue
		}

		invalidLoopCount = 0

		for _, v := range matchTickets {
			userId := uint64(v.SearchFields.DoubleArgs["userId"])
			delete(user2Tickets, userId)
		}

		matches = append(matches, &pb.Match{
			MatchId:       fmt.Sprintf("profile-%v-time-%v-%v", p.GetName(), time.Now().Format("2006-01-02T15:04:05.00"), count),
			MatchProfile:  p.GetName(),
			MatchFunction: matchName,
			Tickets:       matchTickets,
		})

		count++
	}

	return matches, nil
}
