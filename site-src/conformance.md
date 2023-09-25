Adapted with :blue_heart: from the [gateway api project release documentation](https://gateway-api.sigs.k8s.io/)

# Conformance

This API covers a broad set of features and use cases and has been implemented
widely. This combination of both a large feature set and variety of
implementations requires clear conformance definitions and tests to ensure the
API provides a consistent experience wherever it is used.

When considering Network Policy API conformance, there are three important concepts:

## 1. Release Channels

Within Network Policy API, release channels are used to indicate the stability of a
field or resource. The "standard" channel of the API includes fields and
resources that have graduated to "beta". The "experimental" channel of the API
includes everything in the "standard" channel, along with experimental fields
and resources that may still be changed in breaking ways **or removed
altogether**. For more information on this concept, refer to our
[versioning](./versioning.md) documentation.

## 2. Support Levels

Unfortunately some implementations of the API will not be able to support every
feature that has been defined. To address that, the API defines a corresponding
support level for each feature:

* **Core** features will be portable and we expect that there is a reasonable
  roadmap for ALL implementations towards support of APIs in this category.
* **Extended** features are those that are portable but not universally
  supported across implementations. Those implementations that support the
  feature will have the same behavior and semantics. It is expected that some
  number of roadmap features will eventually migrate into the Core. Extended
  features will be part of the API types and schema.
* **Implementation-specific** features are those that are not portable and are
  vendor-specific. Implementation-specific features will not have API types and
  schema except via generic extension points.

Behavior and feature in the Core and Extended set will be defined and validated
via behavior-driven conformance tests. Implementation-specific features will not
be covered by conformance tests.

By including and standardizing Extended features in the API spec, we expect to
be able to converge on portable subsets of the API among implementations without
compromising overall API support. Lack of universal support will not be a
blocker towards developing portable feature sets. Standardizing on spec will
make it easier to eventually graduate to Core when support is widespread.

### Overlapping Support Levels

It is possible for support levels to overlap for a specific field. When this
occurs, the minimum expressed support level should be interpreted. For example,
an identical struct may be embedded in two different places. In one of those
places, the struct is considered to have Core support while the other place only
includes Extended support. Fields within this struct may express separate Core
and Extended support levels, but those levels must not be interpreted as
exceeding the support level of the parent struct they are embedded in.

## 3. Conformance Tests

Network Policy API includes a set of conformance tests. These create a series of
ANPs BANPs while also testing certain features and functionality within each one.

Each release contains a set of conformance tests, these will continue to
expand as the API evolves. Currently conformance tests cover the majority
of Core capabilities in the standard channel, in addition to some Extended
capabilities.

### Running Tests TODO (@tssurya)
