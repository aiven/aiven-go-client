# Aiven

Aiven has a Python Client, but not a Go client. This is a Go SDK which enables
you to use it with Terraform to automatically spin up services together with the
rest of your architecture.

## Development

This SDK is currently still under development. Only the bits that I currently
need are sporadically added. If any new features are required, feel free to open
an issue or a PR.

This has been developed purely for the goal of using this together with
Terraform, but with general usage in mind. You can see the [Terraform provider
here](https://github.com/jelmersnoeck/terraform-provider-aiven)

### Testing

Currently, tests are only run in development and not through CI. This is due to
the requirement of an actual credit card. In the future we'll look into enabling
a mock for this so we can run unit tests on CI and do integration tests on our
own.
