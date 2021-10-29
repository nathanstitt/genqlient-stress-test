this is an example of issues I encountered with using genqlient with a somewhat complex schema

I have a project that's using the Hasura Graphql Engine https://hasura.io which generates a schema for query/mutations on database tables. Hasura has a very comprehensive query language that supports nesting and/or type queries using GraphQL types.

I attempted to migrate to genqlient but ran into several issues with recursive references
and empty fields

Two points of interest are:
 * All the [pointer:true comments](https://github.com/nathanstitt/genqlient-stress-test/blob/main/genqlient.graphql#L49-L170) that I needed for it to compile due to recursive references.  
 * The code in main which [constructs a user object](https://github.com/nathanstitt/genqlient-stress-test/blob/main/main.go#L19-L44), then run it and observe the 
actual structure that's being built (18,938 lines).  The Hasura gql engine completely chokes on it

