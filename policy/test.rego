package policy

import future.keywords.in

# Default deny
default allow := false

# Required predicates that must exist
required_predicates := {"cyclonedx-vex", "testing-results", "promotion"}

# Helper to get all predicates from the evidence connection
existing_predicates := {predicate |
    some edge
    edge = input.data.releaseBundleVersion.getVersion.evidenceConnection.edges[_]
    predicate = edge.node.predicateSlug
}

# Rule to check if all required predicates exist
allow := true if {
    missing_predicates := {x | x = required_predicates[_]; not existing_predicates[x]}
    count(missing_predicates) == 0
}
