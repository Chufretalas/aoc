def get_char_value(char: str):

    if char.isupper():
        return ord(char)-38 # -64 +26

    else:
        return ord(char)-96

def find_wrong_item(sack: str):

    comp1 = sack[:round(len(sack)/2)]
    comp2 = sack[round(len(sack)/2):]

    for char in comp1:
        if comp2.find(char) != -1:
            return char

def main():

    total_points = 0

    with open("./input.txt", "r") as file:
        while True:
            sack = file.readline()
            if sack:
                wrong_item = find_wrong_item(sack.strip())
                points = get_char_value(wrong_item)
                total_points += points
            else:
                break
        file.close()
    
    print(total_points)

if __name__ == "__main__":
    main()