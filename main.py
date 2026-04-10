import sys
import importlib
import time

def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <number>")
        return

    num = sys.argv[1]
    try:
        # Dynamically imports 'solution___.py' and calls its 'solve()' function
        module = importlib.import_module(f"python.solution{num}")
        start_time = time.perf_counter()

        module.solve()

        end_time = time.perf_counter()
        duration = end_time - start_time
        print(f"--- Problem {num} took {duration:.6f}s with Python ---")
    except ImportError:
        print(f"Problem {num} not found.")

if __name__ == "__main__":
    main()
