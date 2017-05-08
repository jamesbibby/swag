package swag

import "github.com/jamesbibby/swag/swagger"

// Builder uses the builder pattern to generate a swagger definition
type Builder struct {
	API *swagger.API
}

// Option provides configuration options to the swagger api builder
type Option func(builder *Builder)

// Description sets info.description
func Description(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Description = v
	}
}

// Version sets info.version
func Version(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Version = v
	}
}

// TermsOfService sets info.termsOfService
func TermsOfService(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.TermsOfService = v
	}
}

// Title sets info.title
func Title(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Title = v
	}
}

// ContactEmail sets info.contact.email
func ContactEmail(v string) Option {
	return func(builder *Builder) {
		builder.API.Info.Contact.Email = v
	}
}

// License sets both info.license.name and info.license.url
func License(name, url string) Option {
	return func(builder *Builder) {
		builder.API.Info.License.Name = name
		builder.API.Info.License.URL = url
	}
}

// BasePath sets basePath
func BasePath(v string) Option {
	return func(builder *Builder) {
		builder.API.BasePath = v
	}
}

// Schemes sets the scheme
func Schemes(v ...string) Option {
	return func(builder *Builder) {
		builder.API.Schemes = v
	}
}

// TagOption provides additional customizations to the #Tag option
type TagOption func(tag *swagger.Tag)

// TagDescription sets externalDocs.description on the tag field
func TagDescription(v string) TagOption {
	return func(t *swagger.Tag) {
		t.Docs.Description = v
	}
}

// TagURL sets externalDocs.url on the tag field
func TagURL(v string) TagOption {
	return func(t *swagger.Tag) {
		t.Docs.URL = v
	}
}

// Tag adds a tag to the swagger api
func Tag(name, description string, options ...TagOption) Option {
	return func(builder *Builder) {
		if builder.API.Tags == nil {
			builder.API.Tags = []swagger.Tag{}
		}

		t := swagger.Tag{
			Name:        name,
			Description: description,
		}

		for _, opt := range options {
			opt(&t)
		}

		builder.API.Tags = append(builder.API.Tags, t)
	}
}

// Host specifies the host field
func Host(v string) Option {
	return func(builder *Builder) {
		builder.API.Host = v
	}
}

// Endpoints allows the endpoints to be added dynamically to the Api
func Endpoints(endpoints ...*swagger.Endpoint) Option {
	return func(builder *Builder) {
		for _, e := range endpoints {
			builder.API.AddEndpoint(e)
		}
	}
}

// New constructs a new api builder
func New(options ...Option) *swagger.API {
	b := &Builder{
		API: &swagger.API{
			BasePath: "/",
			Swagger:  "2.0",
			Schemes:  []string{"http"},
			Info: swagger.Info{
				Contact: swagger.Contact{
					Email: "your-email-address",
				},
				Description:    "Describe your API",
				Title:          "Your API Title",
				Version:        "SNAPSHOT",
				TermsOfService: "http://swagger.io/terms/",
				License: swagger.License{
					Name: "Apache 2.0",
					URL:  "http://www.apache.org/licenses/LICENSE-2.0.html",
				},
			},
		},
	}

	for _, opt := range options {
		opt(b)
	}

	return b.API
}
