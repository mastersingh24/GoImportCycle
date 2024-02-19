# Import Cycle Problem

## Description
This repo is purely for my own learning and memory. When developing multi package repositories it is easy to get to modules that are imported by the same module. But when these modules need to import each other you hit an **import cycle**. 

The main branch of this repo outlines the problem. Three Modules. 

- The Agent
- The Message Server
- The API Server

The Agent implements the API and Message Server as Struct Objects and then kicks off the Agent Which runs the API and Message Server each of which just implment an infinate loop (original aye) outputting their own name. 

The However both the Message Server and the API Server also Try to implement functions of each others... **This creates an import cycle**.

Other Branches in this repo show how to solve the problem using dependency injection and interfaces. 

