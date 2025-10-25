🧠 Golden Test Set – Text Editing & Auto-Correction Tool

This document contains a curated Golden Test Set designed for verifying the correctness and robustness of the text transformation tool written in Go.
It includes standard functional examples (from the audit samples), tricky prototype cases, and a comprehensive “big paragraph” case to ensure coverage of all transformation rules.

🥇 1. Success Test Cases (Basic Functional Examples)

These represent the core transformation rules that every correct implementation must handle.

Case 1 – Hexadecimal Conversion

Input:
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

Expected Output:
Simply add 66 and 2 and you will see the result is 68.

Purpose:
Tests replacement of (hex) and (bin) tokens by converting preceding hexadecimal or binary words into decimal form.

Case 2 – Uppercase, Lowercase, and Capitalization

Input:
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

Expected Output:
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

Purpose:
Covers (up), (low), (cap) and their numbered variants, testing correct case transformations across multiple words.

Case 3 – Article Correction ("a" → "an")

Input:
There is no greater agony than bearing a untold story inside you.

Expected Output:
There is no greater agony than bearing an untold story inside you.

Purpose:
Validates proper article correction when the following word starts with a vowel or the letter 'h'.

Case 4 – Punctuation Spacing

Input:
Punctuation tests are ... kinda boring ,what do you think ?

Expected Output:
Punctuation tests are... kinda boring, what do you think?

Purpose:
Tests punctuation normalization and proper spacing around commas, ellipses, and question marks.

Case 5 – Apostrophe Handling

Input:
I am exactly how they describe me: ' awesome '

Expected Output:
I am exactly how they describe me: 'awesome'

Purpose:
Ensures quotes are attached tightly to the enclosed text with no inner spaces.

🧩 2. Prototype (Tricky) Examples

These examples test edge cases and combined behaviors of the tool.

Prototype 1 – Mixed Case and Numbered Uppercase

Input:
this is so exciting (up, 3) really (low, 2) WEIRD (cap)

Expected Output:
THIS IS SO exciting really Weird

Focus: Multiple chained transformation markers and boundary interactions.

Prototype 2 – Nested Quotation and Punctuation

Input:
He said: ' what a strange world ,isn't it ? ' then walked away...

Expected Output:
He said: 'what a strange world, isn't it?' then walked away...

Focus: Proper placement of apostrophes and punctuation in proximity.

Prototype 3 – Hex, Bin, and Case Combined

Input:
Values 1e (hex), 1111 (bin) (up, 2) must match expectations.

Expected Output:
Values 30, 15 MUST MATCH expectations.

Focus: Arithmetic conversions combined with multi-word case transformation.

Prototype 4 – Article and Punctuation Interaction

Input:
It was a honor to be part of a event ,wasn't it ?

Expected Output:
It was an honor to be part of an event, wasn't it?

Focus: Double article corrections and punctuation normalization.

Prototype 5 – Multiword Quotation and Case Mix

Input:
As Elton John said: ' i am the most well-known homosexual in the world ' (cap, 4)

Expected Output:
As Elton John said: 'I am the most well-known homosexual in the world' I Am The Most Well-known

Focus: Multi-word quotes combined with capitalization over a specified range.

📜 3. Big Paragraph Test

This comprehensive test covers almost all transformation rules in one large input.
It’s ideal for integration and stress testing.

Input:
there (cap) once was a hero named link (cap, 3) ,he carried 1e (hex) rupees and 10 (bin) arrows. he said: ' this is a honor ' before entering a old temple (up, 2) ,where legends said time (low) STOPS (low, 2) ... a mysterious voice whispered: ' welcome ,hero ' and the adventure began !

Expected Output:
There once was a hero named Link, he carried 30 rupees and 2 arrows. He said: 'this is an honor' before entering an old temple, WHERE LEGENDS SAID time stops... a mysterious voice whispered: 'welcome, hero' and the adventure began!

Purpose:
Validates correct sequencing and coexistence of all rule types:

Capitalization and numbered case control

Numeric conversion (hex/bin)

Article correction (a → an)

Punctuation normalization

Quotation spacing

Multi-rule ordering stability