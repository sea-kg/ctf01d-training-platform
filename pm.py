#!/usr/bin/env python3
# Copyright (c) 2024 Evgenii Sopov <mrseakg@gmail.com>

""" Project Manager for ctf01-training-platform """

import sys
import argparse
import libpmlocal


MAIN_PARSER = argparse.ArgumentParser(
    prog='pm.py',
    description='Helper for manage this project',
)
SUBPARSERS = MAIN_PARSER.add_subparsers(
    title='subcommands',
)

SUBCOMMANDS = [
    libpmlocal.CommandCodegenPython(),
    libpmlocal.CommandPyCheck(),
]

for _sc in SUBCOMMANDS:
    _sc.registry(SUBPARSERS)

arguments = MAIN_PARSER.parse_args()
if 'subparser' not in arguments:
    MAIN_PARSER.print_help(sys.stderr)
    sys.exit(1)

COMMAND = arguments.subparser

for _sc in SUBCOMMANDS:
    if _sc.get_command() == COMMAND:
        _sc.execute(arguments)
