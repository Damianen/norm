package main

import (
    "github.com/Damianen/appie-app/handler"
    "fmt"
    "log"
    "net/http"
)

func main() {
    loginHandler := handler.LoginHandler{}
    teamHandler := handler.TeamHandler{}
    aisleHandler := handler.AisleHandler{}
    shiftRaportHandler := handler.ShiftraportHandler{}
    dashboardHandler := handler.DashboardHandler{}
    statsHandler := handler.StatsHandler{}
    planningHandler := handler.PlanningHandler{}

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

    fmt.Println("Server is listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
