import os

render_width = 640
render_height = 480
tile_size = 50

filename = os.path.join(os.getcwd(), "Assets", "Pages", "test2.bin")
square_num_x = render_width // tile_size
square_num_y = render_height // tile_size

middle_x = square_num_x // 2
middle_y = square_num_y // 2

def generate_wall_map(top: bool, bottom: bool, left: bool, right: bool):
    global filename, square_num_x, square_num_y, middle_x, middle_y
    with open(filename, "wb") as file:
        # Include meta data about Teleporters, Enemies numbers and etc.
        # Writing about Teleporters
        file.write(bytes([top, bottom, left, right]))
        
        for y in range(square_num_y):
            for x in range(square_num_x):
                tile_id = 0  # Default interior

                # Top row
                if y == 0:
                    if x == 0:
                        tile_id = 1
                    elif x == square_num_x - 1:
                        tile_id = 2
                    elif top and x == middle_x:
                        tile_id = 0
                    else:
                        tile_id = 5

                # Bottom row
                elif y == square_num_y - 1:
                    if x == 0:
                        tile_id = 4
                    elif x == square_num_x - 1:
                        tile_id = 3
                    elif bottom and x == middle_x:
                        tile_id = 0
                    else:
                        tile_id = 7

                # Middle rows
                else:
                    if x == 0:
                        if left and y == middle_y:
                            tile_id = 0
                        else:
                            tile_id = 6
                    elif x == square_num_x - 1:
                        if right and y == middle_y:
                            tile_id = 0
                        else:
                            tile_id = 8

                file.write(bytes([tile_id]))

generate_wall_map(False, False, True, False)
