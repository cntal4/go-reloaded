🧰 Text Auto-Formatter Tool
📖 Introduction

The goal of this tool is to automatically modify and format text files according to a specific set of linguistic and formatting rules.
Essentially, it functions as an auto-corrector / text-formatter, ensuring consistency and readability across the text.

⚙️ Usage Overview

The program accepts two file name arguments:

Input File → The source file (remains unchanged)

Output File → The formatted version of the input file

The tool reads from the input file, applies all formatting and transformation rules, and writes the corrected output into the output file.

🧩 Formatting & Conversion Rules
1. 🔢 hex Rule — Convert Hexadecimal to Decimal

Replaces the previous hexadecimal word with its decimal equivalent.

✅ Example
"1E(hex) files were added" → "30 files were added"

2. ⚙️ bin Rule — Convert Binary to Decimal

Replaces the previous binary word with its decimal equivalent.

✅ Example
"It has been 10(bin) years" → "It has been 2 years"

3. 🔠 up Rule — Uppercase Conversion

Converts the previous word (or a specified number of words) to UPPERCASE.

✅ Examples
"Ready, set, go (up)!" → "Ready, set, GO!"
"This is so exciting (up, 2)" → "This is SO EXCITING"

4. 🔡 low Rule — Lowercase Conversion

Converts the previous word (or specified number of words) to lowercase.

✅ Examples
"I should stop SHOUTING (low)" → "I should stop shouting"
"I SHOULD STOP SHOUTING (low, 3)" → "I should stop shouting"

5. 🧍‍♂️ cap Rule — Capitalize Words

Converts the previous word (or specified number of words) to Capitalized Case
(first letter uppercase, remaining letters lowercase).

✅ Examples
"Welcome to the Brooklyn bridge (cap)" → "Welcome to the Brooklyn Bridge"
"welcome to the brooklyn bridge (cap, 5)" → "Welcome To The Brooklyn Bridge"

6. ✍️ Punctuation Spacing Rule

Ensures proper spacing around punctuation marks.

🟢 Single Punctuation Marks

Applies to . , ! ? : ;
They should be attached to the previous word and have one space after.

✅ Example
"I was sitting over there ,and then BAMM !!" → "I was sitting over there, and then BAMM!!"

🟣 Grouped Punctuation Marks

Applies to ..., !?, etc.
They should have no space before and one space after.

✅ Example
"I was thinking ... You were right" → "I was thinking... You were right"

7. 🪞 Quotation Rule (' ')

Single quotes must always:

Appear in pairs

Be tightly attached to the words they enclose (no inner spaces)

✅ Examples
"I am exactly how they describe me: ' awesome '" → "I am exactly how they describe me: 'awesome'"
"As Elton John said: ' I am the most well-known homosexual in the world '" → "As Elton John said: 'I am the most well-known homosexual in the world'"

8. 🅰️ Indefinite Article Rule (a → an)

The article "a" should be changed to "an" if the following word begins with:

A vowel (a, e, i, o, u)

The letter h

✅ Example
"There it was. A amazing rock!" → "There it was. An amazing rock!"

✅ Summary

This tool performs automated text correction and formatting by applying:

Category	Rules
Numeric Conversions	hex, bin
Text Case Adjustments	up, low, cap
Punctuation Formatting	Proper spacing and grouping
Quotation Handling	Ensures paired and clean single quotes
Grammar Correction	Contextual article replacement (a → an)

Absolutely ✅ — here’s a **revamped, GitHub-ready Markdown version** of your documentation for the **architecture selection and comparison** between **Pipeline** and **FSM**, matching the exact professional and visually structured style of your previous document.

It includes your comparison table, clean formatting, and a clear **justification** for selecting the **Pipeline architecture** — written in a way that reads well for project documentation or a technical README.

---

# 🏗️ Architecture Selection and Comparison

### 📘 Pipeline vs. FSM (Finite State Machine)

---

## 🧠 Overview

This document outlines the **selection and comparison** between two potential architectures for our tool — the **Pipeline architecture** and the **Finite State Machine (FSM)** architecture.
Both are widely used design paradigms, but each offers different advantages depending on the system’s goals and operational nature.

Our tool performs **sequential text-processing operations**, making it crucial to choose an architecture that offers **simplicity, efficiency, and scalability** for applying multiple transformation rules in order.

---

## ⚖️ Architecture Comparison

| **Aspect**           | **Pipeline Architecture**                                                       | **FSM Architecture**                                                         |
| -------------------- | ------------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| **Purpose**          | Improve throughput by executing multiple operations concurrently.               | Model systems with distinct states and transitions.                          |
| **Structure**        | Composed of multiple sequential processing stages.                              | Composed of a set of states and transitions.                                 |
| **Concurrency**      | High – Multiple stages can work simultaneously.                                 | Limited – Only one state is active at a time.                                |
| **Data Flow**        | Data flows sequentially from one stage to the next.                             | Transitions depend on events and inputs.                                     |
| **State Memory**     | The system doesn’t "remember" past stages; each stage operates on current data. | The system "remembers" its current state, affecting its behavior.            |
| **Latency Impact**   | Increases with more stages (data must pass through all stages).                 | Can have lower latency, depending on transitions.                            |
| **Throughput Focus** | Focuses on throughput by processing multiple pieces of data simultaneously.     | Focuses on system behavior, which may result in less emphasis on throughput. |
| **Example**          | CPU instruction pipelines (fetch, decode, execute).                             | Traffic lights (Green, Yellow, Red states).                                  |
| **Flexibility**      | Works well with systems needing concurrent data processing.                     | Best for systems with discrete state transitions.                            |
| **Complexity**       | Can be complex to handle dependencies and pipeline hazards.                     | Complexity arises in managing many states and transitions.                   |
| **Scaling**          | Easy to scale by adding more stages, but limited by bottlenecks.                | Difficult to scale due to state explosion.                                   |
| **Error Handling**   | Errors must be handled carefully to avoid pipeline stalls or incorrect data.    | Easier to handle errors by transitioning to an error state.                  |

---

## 🧩 Summary

The **Pipeline Architecture** is best suited for tasks that require **high throughput** and **parallel or sequential data processing** of multiple operations — especially when each operation acts independently on data flowing through a series of well-defined stages.

In contrast, the **FSM Architecture** excels in systems where behavior changes based on **discrete events or conditions**, such as robotics, embedded systems, or state-driven logic controllers.

### ✅ When to Use Each

* **Pipeline** → Ideal for data-heavy systems where each stage performs a clear, ordered transformation.
  *Examples: compilers, CPU instruction stages, text processing pipelines.*

* **FSM** → Ideal for reactive systems where transitions depend on external triggers or system states.
  *Examples: vending machines, network protocols, or game AI.*

---

## 🧱 Final Architecture Selection: **Pipeline**

We have chosen the **Pipeline Architecture** for our text-formatting tool.

### 🎯 Reason for Selection

The **Pipeline architecture** aligns perfectly with the operational structure of our tool:

* Each **rule or transformation** (e.g., hex conversion, case change, punctuation formatting) can be treated as a **separate stage** in the pipeline.
* The **input text flows sequentially** through each processing stage, with each stage performing a well-defined, independent operation.
* The design ensures **simplicity, modularity, and scalability** — new rules can easily be added as new stages without redesigning the entire system.
* Unlike FSMs, the tool does not rely on **conditional state transitions**; instead, it performs a **series of deterministic transformations**, making a Pipeline far more efficient and easier to implement.
* This architecture enhances **throughput** and keeps the **code maintainable**, especially as the set of text-processing rules grows.

---

## 🧩 Conclusion

By implementing a **Pipeline architecture**, we ensure that:

* Each rule operates as an independent, reusable module.
* The text transformation process remains **linear, efficient, and easily extensible**.
* The system achieves **high performance** and **clean modular design** without the overhead of managing complex state transitions inherent in FSMs.

In short, the **Pipeline model** provides the right balance between **efficiency**, **clarity**, and **maintainability** for our auto-formatting tool.

---



