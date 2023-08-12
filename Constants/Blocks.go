package Constants

// LogicGateType describes a logic gate type (wow). Use constants in Constants package to specify a logic gate type
type LogicGateType int
type PartTier int

const (
	AndLogicGate  LogicGateType = 0
	OrLogicGate   LogicGateType = 1
	XorLogicGate  LogicGateType = 2
	NandLogicGate LogicGateType = 3
	NorLogicGate  LogicGateType = 4
	NXorLogicGate LogicGateType = 5

	Tier1 PartTier = 1
	Tier2 PartTier = 2
	Tier3 PartTier = 3
	Tier4 PartTier = 4
	Tier5 PartTier = 5
)

// Blocks contains ShapeID-s of few blocks from scrap mechanic
var Blocks = map[string]string{
	// Concrete
	"Concrete Block 1":       "a6c6ce30-dd47-4587-b475-085d55c6a3b4",
	"Concrete Block 2":       "ff234e42-5da4-43cc-8893-940547c97882",
	"Concrete Block 3":       "e281599c-2343-4c86-886e-b2c1444e8810",
	"Concrete Block Cracked": "f5ceb7e3-5576-41d2-82d2-29860cf6e20e",
	"Concrete Block Slab":    "cd0eff89-b693-40ee-bd4c-3500b23df44e",
	// Wood
	"Wood Block 1":     "df953d9c-234f-4ac2-af5e-f0490b223e71",
	"Wood Block 2":     "1897ee42-0291-43e4-9645-8c5a5d310398",
	"Wood Block 3":     "061b5d4b-0a6a-4212-b0ae-9e9681f1cbfb",
	"Wood Block Scrap": "1fc74a28-addb-451a-878d-c3c605d63811",
	// Metal
	"Metal Block 1":        "8aedf6c2-94e1-4506-89d4-a0227c552f1e",
	"Metal Block 2":        "1016cafc-9f6b-40c9-8713-9019d399783f",
	"Metal Block 3":        "c0dfdea5-a39d-433a-b94a-299345a5df46",
	"Metal Block Scrap":    "1f7ac0bb-ad45-4246-9817-59bdf7f7ab39",
	"Metal Block Rusted":   "220b201e-aa40-4995-96c8-e6007af160de",
	"Metal Block Extruded": "25a5ffe7-11b1-4d3e-8d7a-48129cbaf05e",
	"Steel Block Punched":  "ea6864db-bb4f-4a89-b9ec-977849b6713a",
	"Aluminium Block":      "3e3242e4-1791-4f70-8d1d-0ae9ba3ee94c",
	"Metal Block Worn":     "d740a27d-cc0f-4866-9e07-6a5c516ad719",
	// Stone
	"Stone Block Scrap": "30a2288b-e88e-4a92-a916-1edbfc2b2dac",
	// Glass
	"Glass Block":         "5f41af56-df4c-4837-9b3c-10781335757f",
	"Glass Block Tile":    "749f69e0-56c9-488c-adf6-66c58531818f",
	"Glass Block Armored": "b5ee5539-75a2-4fef-873b-ef7c9398b3f5",
	// Plastic
	"Plastic Block":        "628b2d61-5ceb-43e9-8334-a4135566df7a",
	"Plastic Block Bubble": "f406bf6e-9fd5-4aa0-97c1-0b3c2118198e",
	"Plaster Block":        "b145d9ae-4966-4af6-9497-8fca33f9aee3",
	// Net
	"Net Block":          "4aa2a6f0-65a4-42e3-bf96-7dec62570e0b",
	"Net Block Solid":    "3d0b7a6e-5b40-474c-bbaf-efaa54890e6a",
	"Net Block Stripped": "a479066d-4b03-46b5-8437-e99fec3f43ee",
	// Generic
	"Barrier Block":         "09ca2713-28ee-4119-9622-e85490034758",
	"Tile Block":            "8ca49bff-eeef-4b43-abd0-b527a567f1b7",
	"Brick Block":           "0603b36e-0bdb-4828-b90c-ff19abcdfe34",
	"Path Light Block":      "073f92af-f37e-4aff-96b3-d66284d5081c",
	"Spaceship Block":       "027bd4ec-b16d-47d2-8756-e18dc2af3eb6",
	"Cardboard Block":       "f0cba95b-2dc4-4492-8fd9-36546a4cb5aa",
	"Insulation Block":      "9be6047c-3d44-44db-b4b9-9bcf8a9aab20",
	"Carpet Block":          "febce8a6-6c05-4e5d-803b-dfa930286944",
	"Painted Wall Block":    "e981c337-1c8a-449c-8602-1dd990cbba3a",
	"Square Mesh Block":     "b4fa180c-2111-4339-b6fd-aed900b57093",
	"Restroom Block":        "920b40c8-6dfc-42e7-84e1-d7e7e73128f6",
	"Diamond Plate Block":   "f7d4bfed-1093-49b9-be32-394c872a1ef4",
	"Spaceship Floor Block": "4ad97d49-c8a5-47f3-ace3-d56ba3affe50",
	"Sand Block":            "c56700d9-bbe5-4b17-95ed-cef05bd8be1b",

	"Logic Gate": "9f0f56e8-2c31-4d83-996c-d00a9b296c3f",
	"Sensor 5":   "20dcd41c-0a11-4668-9b00-97f278ce21af",
	"Sensor 4":   "de018bc6-1db5-492c-bfec-045e63f9d64b",
	"Sensor 3":   "90fc3603-3544-4254-97ef-ea6723510961",
	"Sensor 2":   "cf46678b-c947-4267-ba85-f66930f5faa4",
	"Sensor 1":   "1d4793af-cb66-4628-804a-9d7404712643",
	"Timer":      "8f7fd0e7-c46e-4944-a414-7ce2437bb30f",
	"Spud Gun":   "1e8d93a4-506b-470d-9ada-9c0a321e2db5",
	"Switch":     "7cf717d7-d167-4f2d-a6e7-6b2c70aa3986",
	"Button":     "0229f59d-1ba3-446a-8f9c-df8b0f816e51",
}

// DefaultColors contains base colors of few blocks from scrap mechanic
var DefaultColors = map[string]string{
	// Concrete
	"Concrete Block 1":       "8d8f89",
	"Concrete Block 2":       "8d8f89",
	"Concrete Block 3":       "c9d7dc",
	"Concrete Block Cracked": "8d8f89",
	"Concrete Block Slab":    "af967b",
	// Wood
	"Wood Block 1":     "9b683a",
	"Wood Block 2":     "dc9153",
	"Wood Block 3":     "f2ad74",
	"Wood Block Scrap": "cd9d71",
	// Metal
	"Metal Block 1":        "675f51",
	"Metal Block 2":        "869499",
	"Metal Block 3":        "88a5ac",
	"Metal Block Scrap":    "df6226",
	"Metal Block Rusted":   "738192",
	"Metal Block Extruded": "858795",
	"Steel Block Punched":  "888888",
	"Aluminium Block":      "727272",
	"Metal Block Worn":     "66837c",
	// Stone
	"Stone Block Scrap": "848484",
	// Glass
	"Glass Block":         "e4f8ff",
	"Glass Block Tile":    "c2f9ff",
	"Glass Block Armored": "3abfb1",
	// Plastic
	"Plastic Block":        "0b9ade",
	"Plastic Block Bubble": "9acfd2",
	"Plaster Block":        "979797",
	// Net
	"Net Block":          "435359",
	"Net Block Solid":    "888888",
	"Net Block Stripped": "888888",
	// Generic
	"Barrier Block":         "ce9e0c",
	"Tile Block":            "bfdfed",
	"Brick Block":           "af967b",
	"Path Light Block":      "727272",
	"Spaceship Block":       "820a0a",
	"Cardboard Block":       "a48052",
	"Insulation Block":      "fff063",
	"Carpet Block":          "368085",
	"Painted Wall Block":    "eeeeee",
	"Square Mesh Block":     "c36512",
	"Restroom Block":        "607b79",
	"Diamond Plate Block":   "43494d",
	"Spaceship Floor Block": "dadada",
	"Sand Block":            "c69146",

	"Logic Gate": "df7f01",
}
