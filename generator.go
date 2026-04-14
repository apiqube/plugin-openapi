package main

// GenerateInput is what the host passes to the exported `generate` function.
type GenerateInput struct {
	SourceURL  string `json:"source_url,omitempty"`  // http://localhost/swagger.json
	SourcePath string `json:"source_path,omitempty"` // local file path
	OutputDir  string `json:"output_dir"`            // where to write generated test files
	Format     string `json:"format,omitempty"`      // full, compact, oneliner
}

// GenerateOutput is what `generate` returns.
type GenerateOutput struct {
	FilesCreated []string `json:"files_created"`
	TestsCount   int      `json:"tests_count"`
	Warnings     []string `json:"warnings,omitempty"`
	Error        string   `json:"error,omitempty"`
}

// generate reads an OpenAPI spec and produces qube test files.
// This is the main function of the generator plugin.
func generate(in *GenerateInput) *GenerateOutput {
	// TODO: implementation
	//
	// 1. Fetch spec:
	//    - From URL via host_http_request
	//    - Or from file via host_read_file
	// 2. Parse OpenAPI JSON/YAML (use kin-openapi compatible structure)
	// 3. Walk paths, group by tag (or by resource prefix)
	// 4. Detect CRUD patterns per resource:
	//    POST /users  →  creates, extract id from response
	//    GET /users/{id}  →  uses saved id
	//    PUT /users/{id}  →  uses saved id, body from schema
	//    DELETE /users/{id}  →  uses saved id
	//    → emit as a scenario file
	// 5. For standalone endpoints: emit as independent tests
	// 6. For each schema property: use fake.* generator
	// 7. For each response code: emit expect.status assertion
	// 8. Write files to OutputDir
	// 9. Return GenerateOutput with file list and counts
	return &GenerateOutput{Error: "not implemented"}
}
