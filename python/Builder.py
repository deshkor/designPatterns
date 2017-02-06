class Player():

    def __init__(self):
        self.name = None
        self.hp = None
        self.damage = None
        self.armor = None
        self.potions = None

    def attack(self):
        print(self.name + "attacks with " + str(self.damage))
        return self.damage

    def defend(self, damage):
        if damage >= self.armor:
            print(self.name + " received " + str((damage - self.armor)) + " damage")
            self.hp -= (damage - self.armor)

        if self.hp <= 0:
            print(self.name + " is dead !")
            return True

        if self.hp <= 10:
            self.heal()

        return False

    def heal(self):
        if self.potions is not None and self.potions >= 0:
            self.potions -= 1
            self.hp += 10
            print(self.name + " used a potion and now has " + str(self.hp) + " hp")

    def getHp(self):
        return self.hp

class PlayerBuilder():
    def __init__(self):
        self.player = Player()

    def setName(self, name):
        self.player.name = name

    def setHp(self, hp):
        self.player.hp = hp

    def setDamage(self, damage):
        self.player.damage = damage

    def setArmor(self, armor):
        self.player.armor = armor

    def setPotions(self, potions):
        self.player.potions = potions

    def getPlayer(self):
        tmp = self.player
        self.player = Player()
        return tmp

if __name__ == "__main__":
    print("Starting game !")

    pb = PlayerBuilder()
    pb.setName("SUPER_HERO")
    pb.setHp(20)
    pb.setDamage(7)
    pb.setArmor(7)
    pb.setPotions(3)
    player = pb.getPlayer()

    enemyFactor = 0.2
    enemy = None

    while True:

        if enemy is None or enemy.getHp() <= 0:
            pb.setName("ENEMY")
            pb.setHp(10 * enemyFactor)
            pb.setDamage(3 * enemyFactor)
            pb.setArmor(2 * enemyFactor)
            pb.setPotions(0)
            enemy = pb.getPlayer()

            enemyFactor += 0.3

        if enemy.defend(player.attack()):
            continue

        if player.defend(enemy.attack()):
            break

    print("Game Over !!!")
