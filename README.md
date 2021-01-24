# common

This is a Go package that contains commoly used tools and struct definitions across multiple microservices.

## acl

This is supposed to contain access control tools that can be used across multiple micoservices.
The main purpose is to restrict access to Discord facing services.

## events

Contains commonly used and generated events by Teeworlds servers.
These evens are created by a monitor that is connected to a Teeworlds server via Econ (telnet) and pushed by the monitor to their corresponding topic of the Mosquitto message broker.

## mqtt

Mosquitto wrappers to easily communicate with the mosquitto docker container.
This is the most used package across all ofthe commonly used tools, as theMosquitto message broker is the central messaging system.
