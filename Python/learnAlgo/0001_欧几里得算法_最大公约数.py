#!/usr/bin/env python3
# -*- coding: utf-8 -*-
__author__ = 'lewin'
__date__ = '2019/5/3'
"""

"""

import os, sys, doctest


# ------------------------- Project Environment -------------------------
def _find_root(n):
    if n > 0: return os.path.dirname(_find_root(n - 1))
    return os.path.abspath(__file__)


_path_project = _find_root(2)
if _path_project not in sys.path: sys.path.insert(0, _path_project)


# ------------------------- Functions -------------------------
def gcd(p: int, q: int) -> int:
    """
    >>> gcd(5,2)
    1
    >>> gcd(4,6)
    2
    >>> gcd(100,25)
    25
    """
    if q == 0:
        return p
    else:
        return gcd(q, p % q)


# ------------------------- Main -------------------------
if __name__ == "__main__":
    doctest.testmod() #没有输出就对了
