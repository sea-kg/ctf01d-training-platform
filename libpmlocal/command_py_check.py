#!/usr/bin/env python3
# Copyright (c) 2024 Evegnii Sopov <mrseakg@gmail.com>

""" CommandPyCheck """

import os
import sys


class CommandPyCheck:
    """ CommandPyCheck """
    def __init__(self):
        self.__command = "py-check"

    def get_command(self):
        """ get_command """
        return self.__command

    def registry(self, subparsers):
        """ registry subcommand """
        py_check = subparsers.add_parser(
            name=self.__command,
            description='Lint python scripts'
        )
        py_check.set_defaults(subparser=self.__command)

    def execute(self, _):
        """ run pycodestyle and pylint """
        print(self.__command)
        check_files = [
            "libpmlocal",
            "pm.py"
        ]
        ret = 0
        for _file in check_files:
            ret_pycodestyle = os.system(
                "python3 -m pycodestyle " + _file + " --max-line-length=100"
            )
            if ret_pycodestyle != 0:
                ret = 1
            ret_pylint = os.system("python3 -m pylint " + _file)
            if ret_pylint != 0:
                ret = 1
        if ret != 0:
            print("FAILED")
        sys.exit(ret)
