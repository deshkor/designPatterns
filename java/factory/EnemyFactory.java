public class EnemyFactory {
	public Enemy getEnemy(String enemyType) {

		if (enemyType.equalsIgnoreCase("WeakEnemy")) {
			return new WeakEnemy();
		} else if (enemyType.equalsIgnoreCase("StrongEnemy")) {
			return new StrongEnemy();
		}

		return null;
	}
}
