package infrastructures

// Sentry's Go SDK enables reporting messages, errors, and panics.
import (
	"fmt"
	"os"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
func InitSentry(Router *gin.Engine) {
	if err := sentry.Init(sentry.ClientOptions{
		/** Sentry automatically assigns you a Data Source Name (DSN) when you
		* create a project to start monitoring events in your app.
		* A DSN tells a Sentry SDK where to send events so the events are associated
		* with the correct project.
		**/
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Sentry initialization succeed 🎉")
	fmt.Println("--------Sentry Initialized 🎉----")

	// attach the handler as one of your middleware
	Router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
		// WaitForDelivery: true,
		// Timeout: 500,
	}))

	Router.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	})
}
