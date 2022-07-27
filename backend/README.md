# KANKAKU Backend

## Development Process

### Required steps

1. Edit GraphQL schemas (`./graph/*.graphqls`)
2. Run `go generate ./...` (Resolvers will be generated)
3. Edit GraphQL resolvers (`./graph/*.resolvers.go`)

### (Optional steps)

- Add dependencies into `./graph/resolver.go`
- Create custom models in `./graph/model`