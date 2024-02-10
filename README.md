# arraydirective

arraydirective verifies that the array of schema files in GraphQL does not have specified directives.

## Installation

```console
go install github.com/nametake/arraydirective/cmd/arraydirective@latest
```

## Usage

```console
arraydirective -schema="schema/**/*.graphql" -query="query/**/*.graphql" -types="ID" -directives="min,max"
```

```
-directives string
      required: comma-separated list of directives
-query string
      pattern of query (default "query/**/*.graphql")
-schema string
      pattern of schema (default "schema/**/*.graphql")
-types string
      comma-separated list of list types
```

If no types are specified, all arrays will be targeted.

## Example

```console
arraydirective -schema="schema/**/*.graphql" -query="query/**/*.graphql" -types="ID" -directives="list"
```

```graphql
directive @list(min: Int, max: Int) on INPUT_FIELD_DEFINITION | ARGUMENT_DEFINITION

type A {
    id: ID!
    name: String!
}

type B {
    name: String!
}

schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

type Query {
    a: A!
    b: B!
    aList(ids: [ID!]! @list(max: 64)): [A!]!
    bList(ids: [ID!]!): [B!]! # want "argument ids of bList has no list directive"

    strList(strs: [String!]!): [String!]!
}

type Mutation {
    createA(input: CreateAInput!): CreateAPayload!
    createB(input: CreateBInput!): CreateBPayload!
}

input CreateAInput {
    aIDs: [ID!]! @list(max: 10)
}

type CreateAPayload {
    a: [A!]!
}

input CreateBInput {
    bIDs: [ID!]! # want "bIDs has list directive"
}

type CreateBPayload {
    b: [B!]!
}
```
