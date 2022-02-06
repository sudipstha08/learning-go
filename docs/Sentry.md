### Sentry :book:
- Often we get asked what the difference is between Sentry and traditional logging. It makes a lot of sense when you're looking at something like a web application, where your errors typically go to a logfile on disk. When you start looking at mobile or desktop applications things start to become a bit more clear.

- Sentry supports every major language, framework, and library.
  
- You can get started for free. Pricing depends on the number of monthly events, transactions, and attachments that you send Sentry

- Sentry doesn’t impact a web site’s performance.Sentry is a listener/handler for errors that asynchronously sends out the error/event to Sentry.io. This is non-blocking. The error/event only goes out if this is an error.
Global handlers have almost no impact as well, as they are native APIs provided by the browsers.

- The best way to think about things is this: logging provides you with a trail of events. Often those events are errors, but many times they're simply informational. Sentry is fundamentally different because we focus on exceptions, or in other words, we capture application crashes.

- To summarize, Sentry works with your application logging infrastructure, often integrating directly. It does not replace the need for those logs, and it's also not a destination for things that aren't actionable errors or crashes.

- Difference between sentry and traditional logging 
  [https://sentry.io/vs/logging/]

- Sentry's Go SDK enables reporting messages, errors, and panics.
  
- Here we will see to integrate gin framework with sentry

- Sentry automatically assigns you a Data Source Name (DSN) when you create a project to start monitoring events in your app.

- A DSN tells a Sentry SDK where to send events so the events are associated with the correct project.

- If this value is not provided, SDKs will try to read it from the SENTRY_DSN environment variable, where applicable. This fallback does not apply in cases like a web browser, where the concept of environment variables does not exist.

- DSNs are safe to keep public because they only allow submission of new events and related event data; they do not allow read access to any information.

- Where to Find Your DSN
If you forget your DSN, view Settings -> Projects -> Client Keys (DSN) in sentry.io.

- The Parts of the DSN
The DSN configures the protocol, public key, server address, and project identifier. It is composed of the following parts:
```sh
{PROTOCOL}://{PUBLIC_KEY}:{SECRET_KEY}@{HOST}{PATH}/{PROJECT_ID}
```
