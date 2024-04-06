package toolchain

// ToolchainOption is a function type that represents an option for configuring
// a Toolchain.
// It takes a pointer to a Toolchain and modifies its properties.
type ToolchainOption func(*Toolchain)
