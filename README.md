# Decorator Pattern Demonstration in Go

## Overview
This repository demonstrates the application of the Decorator design pattern in Go. The project highlights how to extend the behavior of objects dynamically using decorators, showcasing flexibility and best practices in design patterns and unit testing.

## Pattern Description
The Decorator pattern allows behavior to be added to individual objects, dynamically, without affecting the behavior of other objects from the same class. In this project, we have implemented several decorators to demonstrate this pattern effectively.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and usage of components and decorators.
- **pkg/decorator/**: Houses the component interfaces and their implementations.
  - **component.go**: Defines the `Component` interface.
  - **concrete_component.go**: Implements the `ConcreteComponent`.
  - **decorator.go**: Implements a basic `Decorator`.
  - **decorator_a.go**: Implements `DecoratorA`, extending the basic `Decorator`.
  - **decorator_b.go**: Implements `DecoratorB`, extending the basic `Decorator`.
  - **concrete_component_test.go**: Unit tests for `ConcreteComponent`.
  - **decorator_a_test.go**: Unit tests for `DecoratorA`.
  - **decorator_b_test.go**: Unit tests for `DecoratorB`.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Decorator_Pattern.git
cd Go_Decorator_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Example Output
When you run the application, you should see the following output:
```
ConcreteComponent Operation: ConcreteComponent
DecoratorA Operation: DecoratorA(ConcreteComponent)
DecoratorB Operation: DecoratorB(DecoratorA(ConcreteComponent))
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp