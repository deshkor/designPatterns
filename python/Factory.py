class Enemy():
    def __init__(self):
        self.hp = None
        self.damage = None
        self.armor = None

    def attack(self):
        return self.damage

    def defend(self, damage):
        if damage > self.armor:
            self.hp -= (damage - self.armor)

    def getHp(self):
        return self.hp

class WeakEnemy(Enemy):
    def __init__(self):
        print "weak enemy born !"
        self.hp = 10
        self.damage = 5
        self.armor = 3

class StrongEnemy(Enemy):
    def __init__(self):
        print "strong enemy born !"
        self.hp = 50
        self.damage = 10
        self.armor = 10

class Factory:
    def getEnemy(self, enemyType):
        if enemyType.upper() == "WEAKENEMY":
            return WeakEnemy()
        if enemyType.upper() == "STRONGENEMY":
            return StrongEnemy()

class Player:
    def __init__(self):
        print "a player has entered the game"
        self.hp = 35
        self.damage = 15
        self.armor = 8

    def attack(self):
        return self.damage

    def defend(self, damage):
        if damage > self.armor:
            self.hp -= (damage - self.armor)

    def getHp(self):
        return self.hp

    def fight(self, enemy):
        win = True

        while self.hp != 0 or enemy.getHp() != 0:
            enemy.defend(self.attack())

            if enemy.getHp() <= 0:
                print "Enemy is dead !!"
                win = True
                break

            self.defend(enemy.attack())

            if self.getHp() <= 0:
                print "Player is dead !"
                win = False
                break

        return win

if __name__ == "__main__":
    enemyFactory = Factory()
    player = Player()
    we = enemyFactory.getEnemy("weakenemy")
    player.fight(we)

    se = enemyFactory.getEnemy("strongenemy")
    player.fight(se)

    se = enemyFactory.getEnemy("strongenemy")
    player.fight(se)

    print "Game Over !"