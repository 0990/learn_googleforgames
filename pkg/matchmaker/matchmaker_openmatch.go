package matchmaker

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"open-match.dev/open-match/pkg/pb"
	"sync"
	"time"
)

type MakerOpenMatch struct {
	users map[uint64]time.Time

	watchAssign func(*Match)

	args OpenMatchArgs

	backendSvc  pb.BackendServiceClient
	frontendSvc pb.FrontendServiceClient
}

func NewMatchMakerOpenMatch(args OpenMatchArgs) (*MakerOpenMatch, error) {
	p := &MakerOpenMatch{}
	p.args = args
	p.users = make(map[uint64]time.Time)

	conn, err := grpc.Dial(args.BackendEndPoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	p.backendSvc = pb.NewBackendServiceClient(conn)

	conn, err = grpc.Dial(p.args.FrontendEndPoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	p.frontendSvc = pb.NewFrontendServiceClient(conn)

	return p, nil
}

func (p *MakerOpenMatch) Run() {
	go p.run()
}

func (p *MakerOpenMatch) run() {
	for {
		time.Sleep(time.Second)
		p.match()
	}
}

func (p *MakerOpenMatch) Join(userId uint64) error {
	go p.join(userId)
	return nil
}

func (p *MakerOpenMatch) Leave(userId uint64) error {
	return nil
}

func (p *MakerOpenMatch) join(userId uint64) error {
	req := &pb.CreateTicketRequest{
		Ticket: makeTicket(userId),
	}

	resp, err := p.frontendSvc.CreateTicket(context.Background(), req)
	if err != nil {
		return err
	}

	log.Println("Ticket created successfully, id:", resp.Id)
	go deleteOnAssign(p.frontendSvc, resp)
	return nil
}

func makeTicket(userId uint64) *pb.Ticket {
	// Add logic to populate Ticket data and generate Ticket.
	ticket := &pb.Ticket{
		Id:         "",
		Assignment: nil,
		SearchFields: &pb.SearchFields{
			DoubleArgs: map[string]float64{
				"userId": float64(userId),
			},
			StringArgs: nil,
			Tags:       nil,
		},
		Extensions:      nil,
		PersistentField: nil,
		CreateTime:      nil,
	}
	return ticket
}

func deleteOnAssign(fe pb.FrontendServiceClient, t *pb.Ticket) {
	for {
		got, err := fe.GetTicket(context.Background(), &pb.GetTicketRequest{TicketId: t.GetId()})
		if err != nil {
			log.Fatalf("Failed to Get Ticket %v, got %s", t.GetId(), err.Error())
		}

		if got.GetAssignment() != nil {
			log.Printf("Ticket %v got assignment %v", got.GetId(), got.GetAssignment())
			break
		}

		time.Sleep(time.Second * 1)
	}

	_, err := fe.DeleteTicket(context.Background(), &pb.DeleteTicketRequest{TicketId: t.GetId()})
	if err != nil {
		log.Fatalf("Failed to Delete Ticket %v, got %s", t.GetId(), err.Error())
	}
}

func (p *MakerOpenMatch) WatchMatch(f func(*Match)) {
	p.watchAssign = f
}

func (p *MakerOpenMatch) match() {
	matchRooms := p.matchSimple()
	if p.watchAssign != nil {
		for _, v := range matchRooms {
			p.watchAssign(v)
		}
	}
}

func (m *MakerOpenMatch) matchSimple() []*Match {
	result := make([]*Match, 0)

	profiles := generateProfiles()
	var wg sync.WaitGroup
	for _, p := range profiles {
		wg.Add(1)
		go func(wg *sync.WaitGroup, p *pb.MatchProfile) {
			defer wg.Done()
			matches, err := fetch(m.args, m.backendSvc, p)
			if err != nil {
				log.Printf("Failed to fetch matches for profile %v, got %s", p.GetName(), err.Error())
				return
			}

			log.Printf("Generated %v matches for profile %v", len(matches), p.GetName())
			if err := assign(m.backendSvc, matches); err != nil {
				log.Printf("Failed to assign servers to matches, got %s", err.Error())
				return
			}

			for _, match := range matches {
				var mm = NewMatch()
				for _, t := range match.Tickets {
					userId := t.GetSearchFields().DoubleArgs["userId"]
					mm.Users[uint64(userId)] = time.Now()
				}
				result = append(result, mm)
			}
		}(&wg, p)
		wg.Wait()
	}
	return result
}

func generateProfiles() []*pb.MatchProfile {
	var profiles []*pb.MatchProfile

	profiles = append(profiles, &pb.MatchProfile{
		Name: "mode_based_profile",
		Pools: []*pb.Pool{
			{
				Name: "pool_mode",
			},
		},
	},
	)

	return profiles
}

func fetch(args OpenMatchArgs, be pb.BackendServiceClient, p *pb.MatchProfile) ([]*pb.Match, error) {
	req := &pb.FetchMatchesRequest{
		Config: &pb.FunctionConfig{
			Host: args.MatchFunctionHost,
			Port: args.MatchFunctionPort,
			Type: pb.FunctionConfig_GRPC,
		},
		Profile: p,
	}

	stream, err := be.FetchMatches(context.Background(), req)
	if err != nil {
		log.Println()
		return nil, err
	}

	var result []*pb.Match
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		result = append(result, resp.GetMatch())
	}

	return result, nil
}

func assign(be pb.BackendServiceClient, matches []*pb.Match) error {
	for _, match := range matches {
		ticketIDs := []string{}
		for _, t := range match.GetTickets() {
			ticketIDs = append(ticketIDs, t.Id)
		}

		conn := fmt.Sprintf("%d.%d.%d.%d:2222", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		req := &pb.AssignTicketsRequest{
			Assignments: []*pb.AssignmentGroup{
				{
					TicketIds: ticketIDs,
					Assignment: &pb.Assignment{
						Connection: conn,
					},
				},
			},
		}

		if _, err := be.AssignTickets(context.Background(), req); err != nil {
			return fmt.Errorf("AssignTickets failed for match %v, got %w", match.GetMatchId(), err)
		}

		log.Printf("Assigned server %v to match %v", conn, match.GetMatchId())
	}

	return nil
}
