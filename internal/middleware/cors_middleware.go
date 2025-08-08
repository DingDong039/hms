package middleware

import (
    "os"
    "strings"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

// CORS returns a middleware for handling CORS
func CORS() gin.HandlerFunc {
    // Helpers to read env with defaults
    getEnv := func(key, def string) string {
        if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
            return v
        }
        return def
    }

    splitCSV := func(s string) []string {
        parts := strings.Split(s, ",")
        out := make([]string, 0, len(parts))
        for _, p := range parts {
            p = strings.TrimSpace(p)
            if p != "" {
                out = append(out, p)
            }
        }
        return out
    }

    // Read configuration from environment
    originsRaw := getEnv("CORS_ALLOWED_ORIGINS", "*")
    methods := splitCSV(getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS"))
    headers := splitCSV(getEnv("CORS_ALLOWED_HEADERS", "Origin,Content-Type,Accept,Authorization"))
    expose := splitCSV(getEnv("CORS_EXPOSE_HEADERS", "Content-Length"))
    allowCreds := strings.EqualFold(getEnv("CORS_ALLOW_CREDENTIALS", "true"), "true")
    maxAgeStr := getEnv("CORS_MAX_AGE", "12h")
    maxAge, err := time.ParseDuration(maxAgeStr)
    if err != nil {
        maxAge = 12 * time.Hour
    }

    cfg := cors.Config{
        AllowMethods:     methods,
        AllowHeaders:     headers,
        ExposeHeaders:    expose,
        AllowCredentials: allowCreds,
        MaxAge:           maxAge,
    }

    if originsRaw == "*" {
        cfg.AllowAllOrigins = true
    } else {
        cfg.AllowOrigins = splitCSV(originsRaw)
    }

    return cors.New(cfg)
}
