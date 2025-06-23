package middleware

import (
	"backendrest/src/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := parts[1]

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// ✅ token valid แล้ว เอา claims ไปใช้ต่อได้เลย
	c.Locals("user", claims)

	return c.Next()
}

/* var secretKey = []byte("your-secret-key")

// JWTMiddleware ตรวจสอบ JWT ใน Authorization header
func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // ตรวจสอบว่า Authorization header มีค่าเป็น Bearer token หรือไม่
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
            return
        }

        // แยก Bearer token ออกจาก header
        parts := strings.Fields(authHeader)
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
            return
        }
        tokenString := parts[1]

        // ตรวจสอบและ parse JWT token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // ตรวจสอบว่า signing method เป็น HMAC-SHA256 หรือไม่
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // เพิ่มข้อมูลจาก token ลงใน context
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }
        userID := claims["sub"].(string)
        ctx := context.WithValue(r.Context(), "userID", userID)
        r = r.WithContext(ctx)

        // เรียก handler ถัดไป
        next.ServeHTTP(w, r)
    })
}
*/
