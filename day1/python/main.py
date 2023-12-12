from p1 import problem_one
# from p2 import problem_two

def main():
    # with open('input.txt') as day4_input:
    #     print(problem_one(day4_input))
    with open('test.txt') as test:
        lines = test.readlines
        print(problem_one(lines))

main()