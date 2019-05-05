#!/usr/bin/env python3
# -*- coding: utf-8 -*-
__author__ = "lewin"
__date__ = "2019/5/5"
"""

"""

import os, sys


# ------------------------- Project Environment -------------------------
def _find_root(n):
    if n > 0: return os.path.dirname(_find_root(n - 1))
    return os.path.abspath(__file__)


_path_project = _find_root(2)
if _path_project not in sys.path: sys.path.insert(0, _path_project)


# ------------------------- Functions -------------------------
def Q__0001():
    """在__enter__时发生错误，不会执行__exit__."""
    n = 0

    class Me:
        def __enter__(self):
            assert n > 0

        def __exit__(self, exc_type, exc_val, exc_tb):
            n = 1
            print(n)

    try:
        with Me():
            pass
    except:
        pass
    print("如果不会执行exit，那么下面将输出0；如果会执行exit，下面将输出1.")
    print(n)
    assert n==0



# ------------------------- Main -------------------------
if __name__ == "__main__":
    Q__0001()
