package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/teejays/clog"
	gopi "github.com/teejays/gopi"

	"github.com/teejays/goku-util/client/db"

	"github.com/teejays/goku-example-one/backend/gateway"
	http_pharmacy "github.com/teejays/goku-example-one/backend/services/pharmacy/goku.generated/http_handlers"
	"github.com/teejays/goku-example-one/backend/services/users/auth"
	http_users "github.com/teejays/goku-example-one/backend/services/users/goku.generated/http_handlers"
)

func main() {
	if err := mainErr(); err != nil {
		log.Fatalf("Error encountered: %s", err)
	}
}

func mainErr() error {
	var err error
	var ctx = context.Background()

	// Initialize the database
	clog.Warnf("Initializing database connection to %s", "users")
	err = db.InitDatabase(ctx, db.Options{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     db.DEFAULT_POSTGRES_PORT,
		Database: "users",
		User:     os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  "disable",
	})
	if err != nil {
		return fmt.Errorf("Initalizing database: %w", err)
	}

	// Initialize the database
	clog.Warnf("Initializing database connection to %s", "pharmacy")
	err = db.InitDatabase(ctx, db.Options{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     db.DEFAULT_POSTGRES_PORT,
		Database: "pharmacy",
		User:     os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  "disable",
	})
	if err != nil {
		return fmt.Errorf("Initalizing database: %w", err)
	}

	// Initialize the Server
	clog.LogToSyslog = false

	// Get the Routes
	var routes []gopi.Route
	routes = append(routes, http_users.GetUsersRoutes()...)
	routes = append(routes, http_pharmacy.GetPharmacyRoutes()...)

	// Middlewares
	preMiddlewareFuncs := []gopi.MiddlewareFunc{gopi.MiddlewareFunc(gopi.LoggerMiddleware)}
	postMiddlewareFuncs := []gopi.MiddlewareFunc{gopi.SetJSONHeaderMiddleware}
	authMiddlewareFunc, err := auth.GetAuthenticateHTTPMiddleware()
	if err != nil {
		return fmt.Errorf("constructing an AuthenticatorFunc: %w", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		var addr = "0.0.0.0"
		var port = 8080
		clog.Warnf("Starting HTTP server at %s:%d", addr, port)
		err := gopi.StartServer(addr, port, routes, authMiddlewareFunc, preMiddlewareFuncs, postMiddlewareFuncs)
		if err != nil {
			log.Fatalf("HTTP Server Error: %s", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var addr = "0.0.0.0"
		var port = 8081
		clog.Warnf("Starting Gateway server at %s:%d", addr, port)
		err := gateway.StartServer(addr, port, "backend/goku.generated/graphql/schema.generated.graphql")
		if err != nil {
			log.Fatalf("HTTP Server Error: %s", err)
		}
	}()

	wg.Wait()

	// Do Stuff
	// if err = doStuff(); err != nil {
	// 	return err
	// }
	return nil
}

// func doStuff() error {
// 	fmt.Println("Doing stuff...")
// 	ctx := context.Background()

// 	// Fetch medicines
// 	// fmt.Println("Fetching Medicines...")
// 	// medicines, err := pharmacy.ListMedicine(ctx, pharmacy.ListMedicineRequest{})
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Println(medicines)

// 	// Add User
// 	fmt.Println("Insert User...")

// 	conn, err := db.NewConnection("users")
// 	if err != nil {
// 		return err
// 	}

// 	var orgA = organization.Organization{
// 		Name: "Laloo Org",
// 		Type: organization.OrganizationType_TypeA,
// 	}

// 	var orgDal = organizationentitydal.OrganizationEntityDAL{}
// 	org, err := orgDal.InsertOrganization(ctx, conn, orgA)
// 	if err != nil {
// 		return err
// 	}
// 	clog.Infof("Org Created:\n%v", printer.PrettyPrint(org))

// 	var userA = user.User{
// 		OrganizationID: org.ID,
// 		CarModel:       "Toyota",
// 		Contact: dawn.Contact{
// 			Name: dawn.PersonName{
// 				FirstName:     "Joe",
// 				MiddleInitial: "K",
// 				LastName:      "Julliani",
// 			},
// 		},
// 		EmergencyContacts: []dawn.Contact{
// 			{
// 				Name: dawn.PersonName{
// 					FirstName: "Donald",
// 					LastName:  "Trump",
// 				},
// 			},
// 			{
// 				Name: dawn.PersonName{
// 					FirstName: "Steve",
// 					LastName:  "Bannon",
// 				},
// 			},
// 			{ // Duplicate
// 				Name: dawn.PersonName{
// 					FirstName: "Steve",
// 					LastName:  "Bannon",
// 				},
// 			},
// 		},
// 		NickNames: []string{"bablloo", "masoo"},
// 		Logins: []users.Event{
// 			{
// 				HappenedAt: time.Now().Add(-2 * time.Hour),
// 				Type:       users.EventType_Login,
// 			},
// 			{
// 				HappenedAt: time.Now(),
// 				Type:       users.EventType_Logout,
// 			},
// 		},
// 	}

// 	var userDAL = userentitydal.UserEntityDAL{}
// 	user, err := userDAL.InsertUser(ctx, conn, userA)
// 	if err != nil {
// 		return err
// 	}
// 	clog.Infof("User Created:\n%v", printer.PrettyPrint(user))

// 	// Select / List Users
// 	req := userentitydal.ListUsersRequest{
// 		Filter: filters_user.UserFilter{
// 			CarModel: filterlib.NewStringCondition(filterlib.EQUAL, "Toyota"),
// 			HavingEmergencyContacts: &filters_global.ContactFilter{
// 				Name: &filters_global.PersonNameFilter{
// 					FirstName: filterlib.NewStringCondition(filterlib.EQUAL, "Donald"),
// 				},
// 			},
// 		},
// 	}

// 	resp, err := userDAL.ListUsers(ctx, conn, req)
// 	if err != nil {
// 		return err
// 	}
// 	clog.Infof("List Users:\n%c", spew.Sdump(resp))

// 	userId := uuid.MustParse("90ada43c-be40-45ba-8287-2171851050df")
// 	u, err := userDAL.GetUser(ctx, conn, userId)
// 	if err != nil {
// 		return err
// 	}
// 	clog.Infof("Get User:\n%c", spew.Sdump(u))
// 	// // Add Medicine
// 	// fmt.Println("Insert Medicine...")
// 	// var medicine = pharmacy.Medicine{
// 	// 	Name:        "Esomax",
// 	// 	Description: "Used for stomach issues",
// 	// 	CompanyID:   company.ID,
// 	// }
// 	// medicine, err = pharmacy.AddMedicine(ctx, medicine)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Println("Medicine inserted...")
// 	// fmt.Printf("%+v\n", medicine)

// 	// // Fetch medicines
// 	// fmt.Println("Listing Medicines...")
// 	// medicines, err = pharmacy.ListMedicine(ctx, pharmacy.ListMedicineRequest{})
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Printf("%+v\n", medicines)

// 	// // Fetch medicines
// 	// fmt.Println("Getting my Medicine...")
// 	// m, err := pharmacy.GetMedicine(ctx, medicine.ID)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Printf("%+v\n", m)
// 	// fmt.Println("Getting my Medicine's Company...")
// 	// c, err := medicine.GetCompany(ctx)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Printf("%+v\n", c)

// 	return nil
// }
