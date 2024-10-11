# Randoop for Go

Randoop for Go is an automated unit test generation tool for Go, inspired by the original Randoop for Java. It uses feedback-directed random testing to generate sequences of function calls that help expose errors and unexpected behavior in Go programs.

## Overview

This tool aims to help Go developers by automatically generating unit tests for functions and methods in Go packages. It discovers available functions, generates random inputs, and creates meaningful test cases that can be used to increase code coverage and catch bugs.

## Project Roadmap

The development of Randoop for Go is broken down into three phases, each building upon the last. Below is a detailed breakdown of each phase.

### Phase 1: Basic Function Discovery and CLI

In this phase, we establish the foundation of the tool, including function discovery, basic random input generation, and a command-line interface.

**Objectives:**
- Set up the project repository and initialize the Go module.
- Implement reflection-based function discovery to list available functions in a given Go package.
- Build a basic command-line interface (CLI) to allow users to specify the target package for test generation.
- Implement random input generation for simple data types such as `int`, `float64`, `string`, and `bool`.
- Execute the discovered functions with random inputs and print the results.

**Completion Criteria:**
- The tool should be able to discover functions from a specified Go package, generate random inputs for basic types, and execute the functions. Results are printed to the console, and panics are handled gracefully.

---

### Phase 2: Enhanced Test Generation and Execution

In this phase, we extend the functionality to support more complex types and implement logic to save generated tests in Go’s `testing` framework format.

**Objectives:**
- Extend random input generation to handle more complex types, such as slices, structs, and maps.
- Implement stateful test generation, allowing the generation of sequences of method calls on objects.
- Create logic for saving generated test cases in Go’s `testing` format, making them executable with `go test`.
- Add basic feedback mechanisms to detect and handle panics or errors during test case execution, using this feedback to refine test generation.

**Completion Criteria:**
- The tool generates meaningful test cases that are saved to disk as valid Go test files.
- These generated test files can be executed using `go test`, and assertions are included based on the function’s return values.

---

### Phase 3: Feedback-Directed Random Testing and Optimization

In this final phase, we focus on refining the test generation process using runtime feedback, improving performance, and maximizing code coverage.

**Objectives:**
- Implement a feedback mechanism that refines test case generation based on execution results (e.g., avoiding inputs that cause panics).
- Add support for code coverage analysis to ensure that generated test cases maximize code exploration.
- Optimize performance to handle large codebases efficiently, potentially adding concurrency for faster test generation.
- Enhance the CLI with additional options, such as controlling the number of test cases, maximum depth of function call sequences, and more.

**Completion Criteria:**
- The tool can generate comprehensive, adaptive test cases that maximize code coverage and can run efficiently on large Go projects.
- Code coverage metrics are integrated and can be tracked to show how much of the code is tested.

---

## Installation

To install, clone the repository and build the CLI tool:

```bash
git clone https://github.com/yourusername/randoop-for-go.git
cd randoop-for-go
go build ./cmd/randoop
