# Tiny graphics library

**Note: work in progress.** This module isn't ready for use yet.

Create fast graphics on slow SPI-connected displays.

This library is inspired by [LVGL](https://lvgl.io/) and [Fyne](https://fyne.io/) and provides a way to create fast user interfaces on slow SPI connected displays.

## Supported displays

Most dislays from the TinyGo drivers repository can be supported. The main requirement is that they have a way to send raw pixel data to the display (`DrawRGBBitmap8`).

## Code style

To make things fast and to reduce memory consumption, there are a few code style considerations apart from the usual Go conventions:

  * All internal colors are stored as the target pixel data. This is usually RGB565 big-endian.
  * Coordinates are of the `int` type in the public API and when doing integer math, but are stored as `int16`.  
    Rationale: `int16` is usually big enough for all coordinates. However, the main compilation target is 32-bit microcontrollers which are most efficient when working with 32-bit integers. Casting on load/store is usually free: ARM has a sign-extending 16-bit load instruction and stores are simply a truncating 16-bit store. Using `int` in the public API is also more convenient for users of the API.
  * No memory allocations should be needed while writing data to the display.
  * The main package (tinygl) does not care about any particular UI style like Material. Such styles should be implemented separately.

## License

BSD 2-clause license, see LICENSE.txt for details.
