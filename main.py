import sys
import importlib

def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <number>")
        return

    num = sys.argv[1]
    try:
        # Dynamically imports 'solution5.py' and calls its 'solve()' function
        module = importlib.import_module(f"python.solution{num}")
        module.solve()
    except ImportError:
        print(f"Problem {num} not found.")

if __name__ == "__main__":
    main()
