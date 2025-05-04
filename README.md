This project has been made to solve the following challenge:

<h1>Form parser</h1>

```xml
<Form>
	<Field Name="program_language" Type="Enumeration(A,B,C)" Optional="False" FieldType="Select">
		<Caption>Pick your programing language</Caption>
		<Labels>
			<Label Name="A">A(+)</Label>
			<Label Name="B">B</Label>
			<Label Name="C">C (All flavors except C#)</Label>
		</Labels>
	</Field>

	<Section Name="experience" Optional="False">
		<Title>Regarding your experience</Title>
		<Contents>
			<Field Name="other" Type="Text([0,200],Lines:4)" Optional="True" FieldType="TextBox">
				<Caption>Other programming experiences</Caption>
			</Field>
			<Field Name="code_repos" Type="File" Optional="True" FieldType="File">
				<Caption>Upload your code repo's in ZIP.</Caption>
			</Field>
		</Contents>
	</Section>
</From>
```

In the example above, you see the XML representation of a survey form. The form is built by a user using a form builder, which means it can have any combination of fields and sections.

Currently, we need to generate a confirmation PDF rendering the form, including the submitted values.

Outline the logic for rendering the XML schema above with the submission and outputs into a PDF while keeping future expansions in mind:

- Extend logic (e.g., render this schema to an HTML web form, form submission validation, form schema validation).
- Support different input form schema data structures (e.g., FormIO’s JSON).
- Support more input field types.
- Support more form elements in addition to fields/inputs and sections, such as comment boxes.

The solution doesn’t need to compile or be fully implemented but should reflect the following qualities and skills:

- Code style
- Code quality
- Architecture
- Problem-solving

---

To run the project run:

```bash
./bin/run
```

And the server should start locally on port `:8080`

To lint the code run:

```bash
./bin/lint

# Use this if you want the linter to attempt to fix errors itself.
./bin/lint  --fix
```

**Note that I would have liked to write tests, but I did not have the time to do so.** I did try to set up the code in a testable way. I would've used [testify](https://github.com/stretchr/testify) to help with assertions and mocking interfaces.

You can change `./templates/form_template.xml` to get a changed form. Note that you don't have to restart the application to change the template, it should reflect the changes on every HTTP request.

---

I have added the files:

- `./internal/form/json_parser.go`
- `./internal/submission/json_renderer.go`
- `./internal/submission/html_renderer.go`
- `./internal/submission/xml_renderer.go`

Which hold structs which still need to be implemented. I did not find the time to implement those, but I still added them to show that it is easy to support different input/output file types as the challenge suggested.

`./internal/form/json_parser.go` would allow for JSON input form templates (currently only XML is possible).

`./internal/submission/{json,html,xml}_renderer.go` would allow for submission confirmations in different file types (currently only supported in PDF).

I also did not find the time to implement input type validation on a form submission. Currently the validation only checks if all non-optional values are submitted. If not it also doesn't really return a user friendly error. This is also something I still wanted to fix, but did not find the time for it.
