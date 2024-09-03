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
	teamHandler := handler.TeamHandler{db}
	aisleHandler := handler.AisleHandler{db}
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

	http.HandleFunc("/team", teamHandler.HandleTeam)
	http.HandleFunc("/dashboard", dashboardHandler.HandleDashboardShow)
	http.HandleFunc("/aisle", aisleHandler.HandleAisle)
	http.HandleFunc("/planning", planningHandler.HandlePlanningShow)
	http.HandleFunc("/stats", statsHandler.HandleStatsShow)
	http.HandleFunc("/shiftraport", shiftRaportHandler.HandleShiftraportShow)
	http.HandleFunc("/management", managementHandler.HandleManagement)

	http.HandleFunc("/managementInsert", managementHandler.HandleInsertManagementShow)
    http.HandleFunc("/stockerInsert", teamHandler.HandleInsertStockerShow)
    http.HandleFunc("/aisleInsert", aisleHandler.HandleAisleInsertShow)

    http.HandleFunc("/aisleUpdate", aisleHandler.HandleAisleUpdateShow)
    http.HandleFunc("/managementUpdate", managementHandler.HandleManagementUpdateShow)

    http.HandleFunc("/aisleGet", aisleHandler.HandleAisleInfoShow)
    http.HandleFunc("/managementGet", managementHandler.HandleManagementInfoShow)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":3001", nil))

	db.Close()
}
