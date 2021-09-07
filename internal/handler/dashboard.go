package handler

import (
	"github.com/opentracing/opentracing-go"
	"github.com/pete911/opentracing-examples/internal/dashboard"
	"github.com/pete911/opentracing-examples/internal/httputil"
	"log"
	"net/http"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	span, ctx := opentracing.StartSpanFromContext(r.Context(), "GetDashboard")
	defer span.Finish()

	dashboard, err := dashboard.GetDashboard(ctx, "123-456", "789-101")
	if err != nil {
		log.Printf("Cannot find flows: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	httputil.WriteJSONResponse(w, dashboard)
}
