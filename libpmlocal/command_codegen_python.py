#!/usr/bin/env python3
# Copyright (c) 2024 Evegnii Sopov <mrseakg@gmail.com>

""" CommandCodegenPython """


class CommandCodegenPython:
    """ CommandCodegenPython """
    def __init__(self):
        self.__command = "codegen-python"

    def get_command(self):
        """ get_command """
        return self.__command

    def registry(self, subparsers):
        """ registry subcommand """
        codegen_python = subparsers.add_parser(
            name=self.__command,
            description='Generate Client Library'
        )
        codegen_python.set_defaults(subparser=self.__command)

    def execute(self, args):
        """ TODO """
        pass
