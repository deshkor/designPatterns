public class Fighter {

	Fighter() {
		System.out.println("New figher ready to rock !!");
	}

	int damage = 20, armor = 10, hp = 40;

	public void fight(Enemy enemy) {
		while (this.hp > 0 || enemy.getHp() > 0) {
			enemy.defend(this.damage);

			if (enemy.getHp() <= 0)
				break;

			this.hp -= enemy.attack();
		}

		if(this.hp <= 0) {
			System.out.println("HAHAHA you're dead !");
		} else {
			System.out.println("You've won !!!");
		}
	}

	public static void main(String[] args) {
	
		EnemyFactory ef = new EnemyFactory();
		Fighter player = new Fighter();

		Enemy enemy = null;

		if ((enemy = ef.getEnemy("WEAKENEMY")) != null) {
			player.fight(enemy);
		}
		
		if ((enemy = ef.getEnemy("STRONGENEMY")) != null) {
			player.fight(enemy);
		}
		
	}
}
