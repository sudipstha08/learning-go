## Claims Based Authorization

- Claims-based identity is a means of authenticating an end user, application or device to another system in a way that abstracts the entity's specific information while providing data that authorizes it for appropriate and relevant interaction

- This authentication method provides user information automatically so applications do not need to request it of the user and the user doesn't have to provide that information separately for different applications.

- Claims can contain information about the user, roles or permissions, useful to build flexible authorization model. Token contains one or more claims and every claim contains some specific information. The token is digitally signed by token issuer when it’s created, so that it can be verified at the receiver end. Token can also contain additional information e.g. expiry date or id.

- A claim is a statement that one subject, such as a person or organization, makes about itself or another subject. For example, the statement can be about a name, group, buying preference, ethnicity, privilege, association or capability. The subject making the claim or claims is the provider. Claims are packaged into one or more tokens that are then issued by an issuer (provider), commonly known as a security token service (STS).[2]

- Within that claims-based identity framework, a secure token service is responsible for issuing, validating, renewing and cancelling security tokens. 

## REFERENCES
1. https://en.wikipedia.org/wiki/Claims-based_identity
2. https://kariera.future-processing.pl/blog/introduction-to-claims-based-authentication-and-authorization-in-net/