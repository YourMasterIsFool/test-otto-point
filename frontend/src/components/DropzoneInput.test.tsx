// DropzoneInput.test.tsx
import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import DropzoneInput from "./DropzoneInput";
import { vi } from "vitest";

// Mock Icon component
vi.mock("@iconify/react/dist/iconify.js", () => ({
  Icon: ({ icon }: { icon: string }) => (
    <div data-testid="mock-icon">{icon}</div>
  ),
}));

// Define mock before all tests
const mockCreateObjectURL = vi.fn(() => "mock-object-url");

beforeAll(() => {
  // Define if not exists (jsdom doesn't provide it)
  if (!("createObjectURL" in URL)) {
    Object.defineProperty(URL, "createObjectURL", {
      writable: true,
      configurable: true,
      value: mockCreateObjectURL,
    });
  } else {
    vi.spyOn(URL, "createObjectURL").mockImplementation(mockCreateObjectURL);
  }
});

describe("DropzoneInput", () => {
  const onChangeMock = vi.fn();

  beforeEach(() => {
    onChangeMock.mockClear();
    mockCreateObjectURL.mockClear();
  });

  test("renders initial UI elements correctly", () => {
    render(<DropzoneInput onChange={onChangeMock} />);
    expect(screen.getByText("Browser File")).toBeInTheDocument();
    expect(screen.getByText("Drag and Drop files")).toBeInTheDocument();
    expect(screen.getByTestId("mock-icon")).toHaveTextContent(
      "bytesize:download"
    );
  });

  test("accepts image file and updates UI", async () => {
    const { container } = render(<DropzoneInput onChange={onChangeMock} />);
    const input = container.querySelector('input[type="file"]')!;
    const file = new File(["dummy-content"], "test.png", { type: "image/png" });

    fireEvent.change(input, { target: { files: [file] } });

    expect(onChangeMock).toHaveBeenCalledWith("mock-object-url");
    console.log('mock calls', onChangeMock.mock.calls);
    expect(mockCreateObjectURL).toHaveBeenCalledWith(file);
    expect(await screen.findByText("test.png - 13 bytes")).toBeInTheDocument();
  });

  test("rejects non-image files", async () => {
    const { container } = render(<DropzoneInput onChange={onChangeMock} />);
    const input = container.querySelector('input[type="file"]')!;
    const file = new File(["dummy-content"], "test.txt", {
      type: "text/plain",
    });

    fireEvent.change(input, { target: { files: [file] } });

    expect(onChangeMock).not.toHaveBeenCalled();
    expect(screen.queryByText("test.txt")).not.toBeInTheDocument();
  });

  test("handles multiple files selection by keeping first one", async () => {
    const { container } = render(<DropzoneInput onChange={onChangeMock} />);
    const input = container.querySelector('input[type="file"]')!;
    const files = [
      new File(["dummy1"], "test1.png", { type: "image/png" }),
      new File(["dummy2"], "test2.png", { type: "image/png" }),
    ];

    fireEvent.change(input, { target: { files } });

    expect(mockCreateObjectURL).toHaveBeenCalledTimes(1);
    expect(await screen.findByText("test1.png - 6 bytes")).toBeInTheDocument();
    expect(screen.queryByText("test2.png")).not.toBeInTheDocument();
  });
});


console.log("mockCreateObjectURL calls:", mockCreateObjectURL.mock.calls);
console.log("onChangeMock calls:", onChangeMock.mock.calls);