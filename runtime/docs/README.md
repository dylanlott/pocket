# RuntimeMgr

This document outlines the purpose of this module, its components and how they all interact with the other modules.

## Contents

- [Overview](#overview)
- [Components](#components)

### Overview

The `RuntimeMgr`'s purpose is to abstract the runtime so that it's easier to test and reason about various configuration scenarios.

It works like a black-box that takes the current environment/machine and therefore the configuration files, flags supplied to the binary, etc. and returns a structure that can be queried for settings that are relevant for the functioning of the modules and the system as a whole.

### Components

This module includes the following components:

- **Config**

  As the name says, it includes, in the form of properties, module specific configurations.

  It also has a `Base` configuration that is supposed to contain more cross-functional settings that cannot really find place in module-specific "subconfigs" (as another way to define module-specific configurations).

- **Genesis**

  The genesis represents the initial state of the blockchain.

  This allows the binary to start with a specific initial state.

  Similarly to `Config`, these are scoped by module as well and currently we have `Persistence` and `Consensus` specific `GenesisState`s

- **Clock**

  Clock is a drop-in replacement for some of the features offered by the `time` package, it acts as an injectable clock implementation used to provide time manipulation while testing.

  By default, the **real** clock is used and while testing it's possible to override it by using the "option" `WithClock(...)`