package middleware

import (
	"fmt"
	"net/url"
	"regexp"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

//	func XSSMiddleware2(c *fiber.Ctx) error {
//		if c.Method() == fiber.MethodGet || c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPatch || c.Method() == fiber.MethodPut {
//			// Get the request body
//			body := c.Body()
//
//			// Apply XSS filtering using Bluemonday
//			sanitizedBody := bluemonday.StrictPolicy().SanitizeBytes(body)
//
//			// Create a new Fiber context with the sanitized content
//			newCtx := c.Copy().Body(sanitizedBody)
//
//			// Continue to the next middleware or route handler with the updated context
//			return c.Next(newCtx)
//		}
//
//		// Continue to the next middleware or route handler
//		return c.Next()
//	}
func XSSMiddleware(c *fiber.Ctx) error {
	// Specify the allowed characters in query parameters
	allowedCharacters := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

	// Get the entire query string
	queryString := string(c.Request().URI().QueryString())
	// Create a new url.Values to store sanitized values
	sanitizedParams := url.Values{}
	// Parse the query string
	queryParams, err := url.ParseQuery(queryString)
	if err != nil {
		return err
	}
	fmt.Println(queryParams)
	// Iterate through query parameters
	for key, values := range queryParams {
		for _, value := range values {
			// Check if the value contains any disallowed characters
			if !allowedCharacters.MatchString(value) {
				// Sanitize the value by escaping HTML special characters
				sanitizedValue := template.HTMLEscapeString(value)
				// Log or handle the detection of potential XSS attack
				fmt.Printf("XSS attack detected in parameter '%s'. Sanitized value: '%s'\n", key, sanitizedValue)

				// Update the query parameter with the sanitized value in the new url.Values
				sanitizedParams.Add(key, sanitizedValue)
			}
		}
	}
	fmt.Println("SANITIZED")
	fmt.Println(sanitizedParams)
	fmt.Println("ORIGINAL URL")
	fmt.Println(c.OriginalURL())
	c.Request().URI().SetQueryString(sanitizedParams.Encode())
	// Continue to the next middleware or route handler
	return c.Next()
}
