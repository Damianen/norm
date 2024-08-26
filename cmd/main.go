package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Damianen/norm/handler"
	"github.com/Damianen/norm/model"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db := model.DbInit()
	//model.CreateTables(db)

	loginHandler := handler.LoginHandler{db}
	teamHandler := handler.TeamHandler{}
	aisleHandler := handler.AisleHandler{}
	shiftRaportHandler := handler.ShiftraportHandler{}
	dashboardHandler := handler.DashboardHandler{}
	statsHandler := handler.StatsHandler{}
	planningHandler := handler.PlanningHandler{}
	managementHandler := handler.ManagementHandler{db}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", loginHandler.HandleLoginShow)
	http.HandleFunc("/login", loginHandler.HandleLogin)
	http.HandleFunc("/logout", loginHandler.HandleLogout)

	http.HandleFunc("/team", teamHandler.HandleTeamShow)
	http.HandleFunc("/dashboard", dashboardHandler.HandleDashboardShow)
	http.HandleFunc("/aisle", aisleHandler.HandleAisleShow)
	http.HandleFunc("/planning", planningHandler.HandlePlanningShow)
	http.HandleFunc("/stats", statsHandler.HandleStatsShow)
	http.HandleFunc("/shiftraport", shiftRaportHandler.HandleShiftraportShow)
	http.HandleFunc("/management", managementHandler.HandleManagement)

	http.HandleFunc("/managementInsert", managementHandler.HandleInsertManagementShow)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	db.Close()
}
