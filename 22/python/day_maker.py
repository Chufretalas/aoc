import sys
import os
import shutil

def main():
    base_dir = os.path.dirname(__file__)
    new_dir_name = sys.argv[1:]

    if len(new_dir_name) > 0:
        new_dir_name = new_dir_name[0]
    else:
        raise Exception("No name as provided for the directory")
    

    # Creating the directory
    try:
        os.mkdir(os.path.join(base_dir, new_dir_name))
    except FileExistsError as err:
        raise FileExistsError("Directory already exists")

    print(f'Directory created')

    # Making the input files
    f = open(os.path.join(base_dir, new_dir_name, "input.txt"), "x")
    f.close()
    print("input.txt created")
    f = open(os.path.join(base_dir, new_dir_name, "example.txt"), "x")
    f.close()
    print("example.txt created")

    # Copying the python template
    src = os.path.join(base_dir, "template.py")
    dst = os.path.join(base_dir, new_dir_name, "part1.py")
    try:
        shutil.copyfile(src, dst)
    except FileNotFoundError as err:
        shutil.rmtree(os.path.join(base_dir, new_dir_name))
        raise FileNotFoundError('No template.py file found in base directory')
    print(f"python template copied")
    print("done")


if __name__ == "__main__":
    main()