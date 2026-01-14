# TDD Workshop

## Query vs Command

https://martinfowler.com/bliki/CommandQuerySeparation.html

Incoming, sent to self, and Outgoing are defined relative to a given component.
Incoming = `componentA.getSomething()`
Sent to self = componentA modifies itself privately
Outgoing = componentA calls componentB privately

| -            | Query               | Command                                         |
|--------------|---------------------|-------------------------------------------------|
| Incoming     | Assert return value | assert return value of *direct* side-effect (1) |
| Sent to self | ❌                   | ❌                                               |
| Outgoing     | ❌                   | Verify message sent (2)                         |

(1) The direct side effect of set_coins is that the internal state of ATM is changed,
which changes the `retrieve()` return value. We do not test the internal state change by
testing something like `get_coins()`, which would have to be added for the purposes of testing.

(2) Verify THAT the call was made, but not its effects. The effects can be tested as an
incoming call to component B.