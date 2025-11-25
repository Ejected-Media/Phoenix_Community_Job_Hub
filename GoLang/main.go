package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// CompanyStruct matches your JSON data model but ready for the Database
type Company struct {
	Name        string   `firestore:"name"`
	Sector      string   `firestore:"sector"`
	Address     string   `firestore:"address"`
	LocationHub string   `firestore:"location_hub"`
	TargetRoles []string `firestore:"target_roles"`
	EntryPath   string   `firestore:"entry_path"`
	CommuteScore string  `firestore:"commute_score"`
	Notes       string   `firestore:"notes"`
}

func main() {
	// 1. SETUP: Context & Database Connection
	ctx := context.Background()
	
	// IMPORTANT: You must download your "serviceAccountKey.json" from Firebase Console
	// and place it in this same folder.
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v", err)
	}
	defer client.Close()

	fmt.Println("✅ Connected to Firestore! Adding companies...")

	// 2. DATA: Define the new companies (ASML & Oracle)
	newCompanies := []Company{
		{
			Name:        "ASML",
			Sector:      "Semiconductor Mfg",
			Address:     "2625 W Geronimo Pl, Chandler, AZ 85224",
			LocationHub: "Chandler (Price Corridor)",
			TargetRoles: []string{"Lithography Service Engineer", "Install Coordinator"},
			EntryPath:   "Service Engineer Trainee Program",
			CommuteScore: "6/10 (Reverse Commute)",
			Notes:       "Critical supplier to Intel. They make the machines that make chips.",
		},
		{
			Name:        "Oracle Cloud",
			Sector:      "Data Center / Cloud",
			Address:     "US-Phoenix-1 Region (Exact DC Location Confidential)",
			LocationHub: "Phoenix / West Valley",
			TargetRoles: []string{"Data Center Technician", "Cloud Ops Engineer"},
			EntryPath:   "Oracle Cloud Infrastructure (OCI) Certification",
			CommuteScore: "Variable (Security Clearance Often Req)",
			Notes:       "Major cloud region hub. Often co-located with major providers like CyrusOne.",
		},
	}

	// 3. EXECUTE: Loop through and add them to the "employers" collection
	for _, company := range newCompanies {
		_, _, err := client.Collection("employers").Add(ctx, company)
		if err != nil {
			log.Fatalf("Failed adding %s: %v", company.Name, err)
		}
		fmt.Printf("🚀 Successfully added: %s\n", company.Name)
	}

	fmt.Println("🎉 Database Update Complete.")
}
