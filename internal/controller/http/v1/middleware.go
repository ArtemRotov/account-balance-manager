package v1

// const (
// 	ctxKeyRequestID = iota
// )

// func setRequestID(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		id := uuid.New().String()

// 		w.Header().Set("X-Request-ID", id)
// 		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
// 	})
// }

// func logRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		logger := s.logger.WithFields(logrus.Fields{
// 			"remote_addr": r.RemoteAddr,
// 			"request_id":  r.Context().Value(ctxKeyRequestID).(string),
// 		})

// 		logger.Infof("started %s %s", r.Method, r.RequestURI)

// 		start := time.Now()

// 		rw := &responseWriter{w, http.StatusOK}
// 		next.ServeHTTP(rw, r)

// 		logger.Infof("completed with %d %s in %v",
// 			rw.code,
// 			http.StatusText(rw.code),
// 			time.Since(start))
// 	})
// }
