package command

/*
 * This file is part of PhoenixBuilder.

 * PhoenixBuilder is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.

 * Copyright (C) 2021-2025 Bouldev
 */

import (
	"encoding/binary"
	"io"
)

type AddInt16YValue struct {
	Value int16
}

func (_ *AddInt16YValue) ID() uint16 {
	return 22
}

func (_ *AddInt16YValue) Name() string {
	return "AddInt16YValueCommand"
}

func (cmd *AddInt16YValue) Marshal(writer io.Writer) error {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(cmd.Value))
	_, err := writer.Write(buf)
	return err
}

func (cmd *AddInt16YValue) Unmarshal(reader io.Reader) error {
	buf := make([]byte, 2)
	_, err := io.ReadAtLeast(reader, buf, 2)
	if err != nil {
		return err
	}
	cmd.Value = int16(binary.BigEndian.Uint16(buf))
	return nil
}
