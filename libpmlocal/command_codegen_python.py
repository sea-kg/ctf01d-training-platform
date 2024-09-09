#!/usr/bin/env python3
# Copyright (c) 2024 Evegnii Sopov <mrseakg@gmail.com>

""" CommandCodegenPython """

import yaml


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
        with open("./api/openapi.yaml", "rt", encoding="utf-8") as _file:
            _openapi = yaml.safe_load(_file)
            print(_openapi)
            _paths = _openapi["paths"]
            for _path in _paths:
                print(_path)
                # print(_paths[_path])
