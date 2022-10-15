package monster

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type fetchUserDetailUsecaseImpl struct {
	UserId          user.UserId
	UserRepository  user.UserRepository
}

type fetchUserDetailUsecase interface {
	Run() (*user.User, *errors.AppError)
}

func NewFetchUserDetailUsecaseImpl(userId , monsterId model3.MonsterId,
	hunterRepository repository2.HunterRepository,
	monsterRepository repository2.MonsterRepository) attackHunterUseCase {

	return attackHunterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl attackHunterUseCaseImpl) Exec() (*model2.Hunter, *errors2.AppError) {
	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	monsterAttackDamage := service2.CalculateAttackHunterDamage(monster, hunter)
	damagedHunter, attackErr := monster.Attack(hunter, monsterAttackDamage)
	if attackErr.HasErrors() {
		return nil, attackErr
	}
	updatedHunter, updateErr := impl.HunterRepository.Update(damagedHunter)
	if updateErr.HasErrors() {
		return nil, updateErr
	}
	return updatedHunter, nil
}
