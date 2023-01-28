package keeper

import (
	"encoding/binary"

	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPeopleCount get the total number of people
func (k Keeper) GetPeopleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PeopleCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPeopleCount set the total number of people
func (k Keeper) SetPeopleCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PeopleCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPeople appends a people in the store with a new id and update the count
func (k Keeper) AppendPeople(
	ctx sdk.Context,
	people types.People,
) uint64 {
	// Create the people
	count := k.GetPeopleCount(ctx)

	// Set the ID of the appended value
	people.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeopleKey))
	appendedValue := k.cdc.MustMarshal(&people)
	store.Set(GetPeopleIDBytes(people.Id), appendedValue)

	// Update people count
	k.SetPeopleCount(ctx, count+1)

	return count
}

// SetPeople set a specific people in the store
func (k Keeper) SetPeople(ctx sdk.Context, people types.People) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeopleKey))
	b := k.cdc.MustMarshal(&people)
	store.Set(GetPeopleIDBytes(people.Id), b)
}

// GetPeople returns a people from its id
func (k Keeper) GetPeople(ctx sdk.Context, id uint64) (val types.People, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeopleKey))
	b := store.Get(GetPeopleIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePeople removes a people from the store
func (k Keeper) RemovePeople(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeopleKey))
	store.Delete(GetPeopleIDBytes(id))
}

// GetAllPeople returns all people
func (k Keeper) GetAllPeople(ctx sdk.Context) (list []types.People) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeopleKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.People
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPeopleIDBytes returns the byte representation of the ID
func GetPeopleIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPeopleIDFromBytes returns ID in uint64 format from a byte array
func GetPeopleIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
