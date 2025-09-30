# Contributor Guide

## Style Guide
- Keep functions small and focused.
- Use context.Context as the first parameter when needed.
- Name clearly without leaving any comments above lines

## Project Structure
Use flat structure in this project, so there is only one package in this project (package main)

## Testing Instructions
- Fix any test or type errors until the whole suite is green.
- Add or update tests for the code you change, even if nobody asked.
- Use table-driven tests.
- Name test functions clearly: TestAddUser_WhenEmailExists_ReturnsError
- Use t.Helper() for helper functions in tests.
- Use only testing package if you want to create mock from interfaces use this mockgen from Uber (go.uber.org/mock/mockgen)

## PR instructions
Title format: [<package_name>] <Short Description>

For example:
[auth] Add JWT token renewal logic

## Working with the `gh` CLI
- Always authenticate with `gh auth login` before running any other commands.
- Use `gh pr checkout <number>` to review pull requests locally.
- Run `gh pr create --fill` to open new pull requests using the prepared commit information.
- Use `gh issue list` and `gh issue view <number>` to triage or review existing issues.
- When in doubt about available commands, run `gh help` or `gh <command> --help` for detailed guidance.
