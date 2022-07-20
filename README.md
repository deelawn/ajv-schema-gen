# ajv-schema-gen
Generates basic AJV JSON schema based on JSON input

## Purpose

The purpose of this schema generation tool is to be able to quickly define schemas for adding tests to Postman.
Note the following constraints:
- For array types, schema is only generated for the first instance in the array
- Not all objects are the same; the JSON body you provide may not contain the entire schema, so check this beforehand

This is far from complete so make an issue or submit a PR for issues.

## Installation and Usage

First make sure you have go installed and then run:
`go install github.com/deelawn/ajv-schema-gen/cmd/ajvgen`

Then copy a JSON response to your clipboard and run the program as follows, saving the result to your clipboard:
`pbpaste | ajvgen | pbcopy`

Inside a postman test start to define a schema:
`var schema = ` and paste after the equals sign to finish defining the schema with `ajv` output.

Then add code to run the schema comparison test
```
pm.test("validating schema", function() {
    pm.response.to.have.jsonSchema(schema);
});
```