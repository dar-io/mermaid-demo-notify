// Package api exposes the client SDK's HTTP surface.
package api

import "net/http"

// Routes is the SDK's public contract.
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("/v1/notifications", listNotifications)
	mux.HandleFunc("/v1/notifications/send", sendNotification)
	mux.HandleFunc("/v1/preferences", updatePreferences)
}

func listNotifications(w http.ResponseWriter, r *http.Request) {}
func sendNotification(w http.ResponseWriter, r *http.Request)  {}
func updatePreferences(w http.ResponseWriter, r *http.Request) {}
