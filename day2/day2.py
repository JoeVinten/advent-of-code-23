import re

color_quantities = {'blue': 14, 'red': 12, 'green': 13}

def sum_to_the_power(game_subsets):
    min_colors = {'blue': 0, 'red': 0, 'green': 0}

    for balls in game_subsets:
        for color in color_quantities:
            pattern = re.compile(fr'(\d+)\s+{color}')
            match = pattern.search(balls)
            if match:
                if min_colors[color] < int(match.group(1)):
                    min_colors[color] = int(match.group(1))
    
    return min_colors['blue'] * min_colors['red'] * min_colors['green']
      

def validate_game(game_name, game_subsets):
    print(f' {game_name} '.center(80, '*'))
    incorrect = False

    for balls in game_subsets:
        for color in color_quantities:
            pattern = re.compile(fr'(\d+)\s+{color}')
            match = pattern.search(balls)

            if match:
                balls_value = int(match.group(1))
                if color_quantities[color] - balls_value < 0:
                    print('this is an incorrect game!')
                    incorrect = True

    return not incorrect 

    

with open('data.txt') as game_data:
    possible_games_total = 0
    total_sum_to_the_power = 0
    for game_line in game_data:
        game_parts = game_line.strip().split(':')
        game_name = game_parts[0]
        game_subsets = game_parts[1].split(';')

        total_sum_to_the_power += sum_to_the_power(game_subsets)

        if validate_game(game_name, game_subsets):
            possible_games_total += int(game_name.split(" ")[1])

print(f"Sum of IDs of possible games: {possible_games_total}")
print(f"Sum of the total powers: {total_sum_to_the_power}")
