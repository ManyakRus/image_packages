Golang module dependency graph, with directories.

The image_packages console utility is designed to display all packages (modules) of any repository in the Golang language
in the form of a scheme diagram(graph) in .graphml format, which can be converted to a .jpg image, etc.
It is necessary to understand the structure of the source code of the repository,
to study or better understand the call structure of the source code.
Displayed:
- directory names
- package names
- number of functions and lines in the package
- go goroutine call arrows (blue)
- function call arrows (black)
- package relationship arrows (dashed)

Sample execution (pictures) can be found in the examples directory

Installation order:
1. Install the .graphml file editor yEd (free)
https://www.yworks.com/products/yed/download

2. Compile this repository
make build
image_packages file will appear in the bin folder

3. Run the image_packages file with parameters:
image_packages <your repository directory> <filename.graphml>

4. Open the resulting .graphml file in the yEd editor
(all elements will be in the center of the screen first)
and select from the menu:
Layout - BPMN
- The yEd editor will arrange all the elements of the circuit in an optimal way.
You can also select a different Layout type to change the display.

5. Export the scheme to a picture.
Select from the menu:
File-Export

![packages](https://github.com/ManyakRus/image_packages/assets/30662875/e56ca425-7fe3-4128-b4d5-2341106ffd6e)



Source code in Golang language.
Tested on Linux Ubuntu
Readme from 10.07.2023

Made by Aleksandr Nikitin
https://github.com/ManyakRus/image_packages
