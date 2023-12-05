# Tiny graphics library

**Note: work in progress.** This module isn't ready for use yet.

Create fast graphics on slow SPI-connected displays.

This library is inspired by [LVGL](https://lvgl.io/) and [Fyne](https://fyne.io/) and provides a way to create fast user interfaces on slow SPI connected displays.

## Supported displays

Most dislays from the TinyGo drivers repository can be supported. The main requirement is that they have a way to send raw pixel data to the display (`DrawRGBBitmap8`).

## Contributing

You are free to contribute, but note that the design isn't complete yet so more features may need to be delayed a bit until I think the design is right.

### Code style

To make things fast and to reduce memory consumption, there are a few code style considerations apart from the usual Go conventions:

  * All internal colors are stored as the target pixel data. This is usually RGB565 big-endian.
  * Coordinates are of the `int` type in the public API and when doing integer math, but are stored as `int16`.  
    Rationale: `int16` is usually big enough for all coordinates. However, the main compilation target is 32-bit microcontrollers which are most efficient when working with 32-bit integers. Casting on load/store is usually free: ARM has a sign-extending 16-bit load instruction and stores are simply a truncating 16-bit store. Using `int` in the public API is also more convenient for users of the API.
  * No memory allocations should be needed while writing data to the display.
  * The main package (tinygl) does not care about any particular UI style like Material. Such styles should be implemented separately.

### Release criteria

This library isn't complete. There are some things I'd like to improve before calling this stable:

  - [x] ~The `Displayer` interface isn't great: `DrawRGBBitmap8` always takes a byte slice, which is unfortunate. It would be better if it could take a slice of the underlying pixel data instead.~
  - [ ] Black-and-white screens aren't well supported: they have pixel data smaller than a single byte. These screens aren't natively supported in LVGL either, but it would be nice if we were able to.
  - [ ] Displays with in-memory buffers could be better supported, by writing to the buffers directly. An example is the hub75 display driver.
  - [ ] Theming needs to be improved. Ideally, the theme and the layout code are entirely separate and the theme just sets the sizes/colors to be used for standard widgets.
  - [ ] DPI scaling isn't implemented yet. It should ideally be able to do all important calculations at compile time.
  - Lots of features are missing, like:
    - [ ] show/hide animations that look smooth (by only redrawing parts of the screen that changed)
    - [x] ~using hardware scrolling present in most SPI displays~
    - [ ] all the missing widgets and container types

## License

BSD 2-clause license, see LICENSE.txt for details.
